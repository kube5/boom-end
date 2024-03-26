package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	derrors "github.com/Mu-Exchange/Mu-End/common/errors"
	pb "github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Mu-Exchange/Mu-End/common/repo"
	"github.com/Mu-Exchange/Mu-End/common/utils/basic"
	"github.com/Mu-Exchange/Mu-End/common/utils/context"
	"github.com/Mu-Exchange/Mu-End/common/utils/uuid"
)

const (
	requestContextKey = "request_context"
)

type microService struct {
	userService pb.UserService
	gameService pb.GameService
}

type Server struct {
	microService
	repo.CommonComponents
	router *gin.Engine
	hs     *http.Server
	tc     TokenCache
}

func New(
	lc fx.Lifecycle,
	sd fx.Shutdowner,
	commonComponents repo.CommonComponents,
	userService pb.UserService,
	gameService pb.GameService,
	tc TokenCache) *Server {
	router := gin.New()
	s := &Server{
		CommonComponents: commonComponents,
		microService: microService{
			userService: userService,
			gameService: gameService,
		},
		tc:     tc,
		router: router,
		hs: &http.Server{
			Addr:           fmt.Sprintf(":%d", commonComponents.Cfg.HTTP.Port),
			Handler:        router,
			ReadTimeout:    commonComponents.Cfg.HTTP.ReadTimeout.Duration,
			WriteTimeout:   commonComponents.Cfg.HTTP.WriteTimeout.Duration,
			MaxHeaderBytes: 1 << 20,
		},
	}
	basic.RegisterHook(lc, sd, s)
	return s
}

func (s *Server) Start(sd fx.Shutdowner) error {
	s.router.MaxMultipartMemory = s.Cfg.HTTP.MultipartMemory
	gin.SetMode(gin.ReleaseMode)
	s.router.Use(s.loggerMiddleware)
	s.router.Use(s.crossOriginMiddleware)
	if err := s.init(); err != nil {
		return errors.Wrap(err, "register router failed")
	}

	printServerInfo := func() {
		s.Logger.Infof("HTTP server listen on : %d", s.Cfg.HTTP.Port)
	}

	go func() {
		err := func() error {
			if s.Cfg.HTTP.TLSEnable {
				if _, err := os.Stat(s.Cfg.HTTP.TLSCertFilePath); err != nil {
					return errors.Wrapf(err, "tls_cert_file_path [%s] is invalid path", s.Cfg.HTTP.TLSCertFilePath)
				}
				if _, err := os.Stat(s.Cfg.HTTP.TLSKeyFilePath); err != nil {
					return errors.Wrapf(err, "tls_key_file_path [%s] is invalid path", s.Cfg.HTTP.TLSKeyFilePath)
				}
				printServerInfo()
				if err := s.hs.ListenAndServeTLS(s.Cfg.HTTP.TLSCertFilePath, s.Cfg.HTTP.TLSKeyFilePath); err != nil {
					return err
				}
			} else {
				printServerInfo()
				if err := s.hs.ListenAndServe(); err != nil {
					return err
				}
			}

			return nil
		}()
		if err != nil {
			s.Logger.Errorf("Server shutdown: %v", err)
		}

		if err := sd.Shutdown(); err != nil {
			s.Logger.Errorf("App shutdown error: %v", err)
		}
	}()

	return nil
}

func (s *Server) Stop(sd fx.Shutdowner) error {
	return s.hs.Close()
}

func (s *Server) init() error {
	v := s.router.Group("/api/v1")
	{

		ug := v.Group("/user")
		ug.GET("/profile", s.apiHandlerWrap(s.Profile))

		ug.POST("/wallet_login_pre", s.apiHandlerWrap(s.MetaMaskLoginPre))
		ug.POST("/wallet_login", s.apiHandlerWrap(s.MetaMaskLogin))
		ug.POST("/refresh_token", s.apiHandlerWrap(s.RefreshToken))
		ug.POST("/logout", s.TokenAuthMiddleware(), s.apiHandlerWrap(s.Logout))

		mg := v.Group("/mission")
		mg.GET("/profile", s.TokenAuthMiddleware(), s.apiHandlerWrap(s.MissionProfile))

		gg := v.Group("/game")
		gg.POST("/action/random", s.TokenAuthMiddleware(), s.apiHandlerWrap(s.GameRandom))

		lg := v.Group("/leaderboard")
		lg.GET("/", s.apiHandlerWrap(s.LeaderBoard))

	}

	t := v.Group("/tg")
	{
		ug := t.Group("/user")
		ug.GET("/profile", s.apiHandlerWrap(s.TGProfile))

		ug.POST("/tg_login_pre", s.apiHandlerWrap(s.TGMetaMaskLoginPre))
		ug.POST("/tg_login", s.apiHandlerWrap(s.TGMetaMaskLogin))
		ug.POST("/refresh_token", s.apiHandlerWrap(s.TGRefreshToken))
		ug.POST("/logout", s.TGTokenAuthMiddleware(), s.apiHandlerWrap(s.TGLogout))
		//ug.POST("/bind/code", s.TGTokenAuthMiddleware(), s.apiHandlerWrap(s.BindCode))

		mg := v.Group("/mission")
		mg.GET("/profile", s.TGTokenAuthMiddleware(), s.apiHandlerWrap(s.TGMissionProfile))

		gg := t.Group("/game")
		gg.POST("/action/random", s.TGTokenAuthMiddleware(), s.apiHandlerWrap(s.TGGameRandom))
		gg.POST("/action/buy/dice", s.TGTokenAuthMiddleware(), s.apiHandlerWrap(s.BuyDice))
		gg.POST("/action/buy/vip", s.TGTokenAuthMiddleware(), s.apiHandlerWrap(s.BuyVip))
		gg.GET("/24h", s.TGTokenAuthMiddleware(), s.apiHandlerWrap(s.Game24H))

		lg := t.Group("/leaderboard")
		lg.GET("/", s.apiHandlerWrap(s.TGLeaderBoard))
	}

	i := s.router.Group("/api/internal")
	{
		i.GET("/wallet_login", s.apiHandlerWrap(s.LoginInternal))
		i.GET("/wallet_login2", s.apiHandlerWrap(s.TGLoginInternal))
	}

	// debug enable pprof
	if gin.IsDebugging() {
		s.router.GET("/debug/pprof/", IndexHandler())
		s.router.GET("/debug/pprof/heap", HeapHandler())
		s.router.GET("/debug/pprof/goroutine", GoroutineHandler())
		s.router.GET("/debug/pprof/allocs", AllocsHandler())
		s.router.GET("/debug/pprof/block", BlockHandler())
		s.router.GET("/debug/pprof/threadcreate", ThreadCreateHandler())
		s.router.GET("/debug/pprof/cmdline", CmdlineHandler())
		s.router.GET("/debug/pprof/profile", ProfileHandler())
		s.router.GET("/debug/pprof/symbol", SymbolHandler())
		s.router.GET("/debug/pprof/trace", TraceHandler())
		s.router.GET("/debug/pprof/mutex", MutexHandler())
	}

	return nil
}

func (s *Server) crossOriginMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

func (s *Server) loggerMiddleware(c *gin.Context) {
	ctx := s.generateRequestContext(c)

	startTime := time.Now()
	c.Next()
	endTime := time.Now()
	if c.Request.Method == "OPTIONS" {
		return
	}

	latencyTime := fmt.Sprintf("%6v", endTime.Sub(startTime))
	reqMethod := c.Request.Method
	reqUri := c.Request.RequestURI
	statusCode := c.Writer.Status()
	clientIP := c.ClientIP()

	loggerFields := logrus.Fields{
		"http_code":  statusCode,
		"total_time": latencyTime,
		"ip":         clientIP,
		"method":     reqMethod,
		"uri":        reqUri,
	}

	if ctx != nil {
		for key, value := range ctx.LogFields() {
			loggerFields[key] = value
		}
	}

	ctx.Logger.WithFields(loggerFields).Info("Handle request")
}

func (s *Server) generateRequestContext(c *gin.Context) *context.Context {
	requestIDStr := uuid.GenUUID().Encode()
	ctx := context.NewContext(requestIDStr, s.Logger)
	c.Set(requestContextKey, ctx)

	return ctx
}

type Handler func(ctx *context.Context, c *gin.Context) (res interface{}, err error)
type HandlersChain []Handler

func (s *Server) apiHandlerWrap(handler ...Handler) func(c *gin.Context) {
	return s.handlerWrapper(0, handler)
}

func (s *Server) handlerWrapper(successCode int, handler HandlersChain) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := s.getRequestContext(c)
		var res interface{}
		err := func() (gErr error) {
			defer func() {
				if err := recover(); err != nil {
					gErr = fmt.Errorf("%v", err)
				}
			}()
			res, gErr = handler[0](ctx, c)
			return gErr
		}()
		if err != nil {
			s.failResponseWithErr(c, err)
			return
		}

		for i := 1; i < len(handler); i++ {
			_, err := handler[i](ctx, c)
			if err != nil {
				ctx.Logger.WithFields(logrus.Fields{
					"handle_method": reflect.TypeOf(handler[i]).Name(),
					"err_msg":       err.Error(),
				}).Error("Handle Wrapper Err!")
			}
		}
		s.successResponseWithData(c, successCode, res)
	}
}

func (s *Server) successResponse() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}

func (s *Server) getRequestContext(c *gin.Context) *context.Context {
	ctxValue, ok := c.Get(requestContextKey)
	if ok {
		return ctxValue.(*context.Context)
	}
	return context.NewContext("", s.Logger)
}

func (s *Server) failResponseWithErr(c *gin.Context, err error) {
	code := derrors.DecodeError(errors.New(err.Error()))
	msg := err.Error()

	ctx := s.getRequestContext(c)
	ctx.AddLogField("err_code", code)
	ctx.AddLogField("err_msg", msg)
	if cb, ok := c.Get(gin.BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			ctx.AddLogField("params", string(cbb))
		}
	}

	if token, err := s.ExtractTokenMetadata(c); err == nil {
		ctx.AddLogField("user_id", token.UserId)
		ctx.AddLogField("access_uuid", token.AccessUuid)
	}
	req := ctx.Req()
	// log request on failure
	if req != nil {
		reqStr, err := json.Marshal(req)
		if err != nil {
			ctx.AddLogField("req", string(reqStr))
		}
	}
	httpCode := http.StatusOK
	c.JSON(httpCode, gin.H{
		"code":    code,
		"message": ternary(code == derrors.UnknownErrorCode, derrors.UnknownErrorMsg, msg),
	})
}

func (s *Server) successResponseWithData(c *gin.Context, successCode int, data interface{}) {
	res := gin.H{
		"code": successCode,
	}
	if data != nil {
		res["data"] = data
	}
	c.JSON(http.StatusOK, res)
}

func ternary(a bool, b, c interface{}) interface{} {
	if a {
		return b
	}
	return c
}
