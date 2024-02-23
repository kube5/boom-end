package pyth

import (
	"context"
	"testing"
)

func TestEvmPriceServiceConnection_GetPriceFeedsUpdateData(t *testing.T) {
	type fields struct {
		url string
		ctx context.Context
	}
	type args struct {
		priceIds []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				url: "https://mock-pyth.rollie.finance",
				ctx: context.Background(),
			},
			args: args{
				priceIds: []string{
					"0xff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace", // ETH
					"0xe62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b43", // BTC
				},
			},
			want: []string{
				"0xff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace0000000000000000000000000000000000000000000000000000003339059800", // 2200
				"0xe62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b4300000000000000000000000000000000000000000000000000000417bce6c800", // 45000
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &EvmPriceServiceConnection{
				url: tt.fields.url,
				ctx: tt.fields.ctx,
			}
			got, err := c.GetPriceFeedsUpdateData(tt.args.priceIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPriceFeedsUpdateData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("GetPriceFeedsUpdateData() got = %v, want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("GetPriceFeedsUpdateData() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
