pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract Staking is ReentrancyGuard{

    IERC20 public immutable eth;
    IERC20 public immutable usdt;
    IERC20 public immutable usdc;

    uint256 public totalStakedEth;
    uint256 public totalStakedUsdt;
    uint256 public totalStakedUsdc;

    struct Staker {
        uint256 ethStaked;
        uint256 usdtStaked;
        uint256 usdcStaked;
        bool isMintDice;
    }

    mapping(address => Staker) public stakers;

    constructor(
        address _eth,
        address _usdt,
        address _usdc
    ) {
        eth = IERC20(_eth);
        usdt = IERC20(_usdt);
        usdc = IERC20(_usdc);
    }

    event MintDice(address indexed staker, address token);
    event StakeEth(address indexed staker, uint256 amount);
    event StakeUsdt(address indexed staker, uint256 amount);
    event StakeUsdc(address indexed staker, uint256 amount);
    event Withdraw(address indexed staker, address token ,uint256 amount);

    function mintDice(address token) external nonReentrant{
        Staker storage staker = stakers[msg.sender];
        require(staker.isMintDice == false, "Invalid has been mint dice");
        require(token == address(usdt) || token == address(usdc), "Invalid token");
        uint8 decimals = getDecimals(token);
        if (token == address(usdt)) {
            require(IERC20(token).balanceOf(msg.sender) >= 10**decimals * 10, "Insufficient balance");
            usdt.transferFrom(msg.sender, address(this), 10**decimals * 10);
        } else {
            require(IERC20(token).balanceOf(msg.sender) >= 10**decimals * 10, "Insufficient balance");
            usdc.transferFrom(msg.sender, address(this), 10**decimals * 10);
        }
        staker.isMintDice = true;
        emit MintDice(msg.sender, token);
    }

    function stakeEth(uint256 amount) external nonReentrant{
        require(amount > 0, "Invalid amount");
        require(eth.balanceOf(msg.sender) >= amount, "Insufficient balance");

        eth.transferFrom(msg.sender, address(this), amount);

        Staker storage staker = stakers[msg.sender];
        staker.ethStaked += amount;
        totalStakedEth += amount;
        emit StakeEth(msg.sender, amount);
    }

    function stakeUsdt(uint256 amount) external nonReentrant{
        require(amount > 0, "Invalid amount");
        require(usdt.balanceOf(msg.sender) >= amount, "Insufficient balance");

        usdt.transferFrom(msg.sender, address(this), amount);

        Staker storage staker = stakers[msg.sender];
        staker.usdtStaked += amount;
        totalStakedUsdt += amount;
        emit StakeUsdt(msg.sender, amount);
    }

    function stakeUsdc(uint256 amount) external nonReentrant{
        require(amount > 0, "Invalid amount");
        require(usdc.balanceOf(msg.sender) >= amount, "Insufficient balance");

        usdc.transferFrom(msg.sender, address(this), amount);

        Staker storage staker = stakers[msg.sender];
        staker.usdcStaked += amount;
        totalStakedUsdc += amount;
        emit StakeUsdc(msg.sender, amount);
    }

    function withdraw(uint256 amount, address token) external nonReentrant{
        require(amount > 0, "Invalid amount");
        require(token == address(eth) || token == address(usdt) || token == address(usdc), "Invalid token");

        Staker storage staker = stakers[msg.sender];

        if (token == address(eth)) {
            require(staker.ethStaked >= amount, "Insufficient staked ETH");
            eth.transfer(msg.sender, amount);
            staker.ethStaked -= amount;
            totalStakedEth -= amount;
        } else if (token == address(usdt)) {
            require(staker.usdtStaked >= amount, "Insufficient staked USDT");
            usdt.transfer(msg.sender, amount);
            staker.usdtStaked -= amount;
            totalStakedUsdt -= amount;
        } else {
            require(staker.usdcStaked >= amount, "Insufficient staked USDC");
            usdc.transfer(msg.sender, amount);
            staker.usdcStaked -= amount;
            totalStakedUsdc -= amount;
        }
        emit Withdraw(msg.sender, token, amount);
    }

    function getStakedTotals() public view returns (uint256, uint256, uint256) {
        return (totalStakedEth, totalStakedUsdt, totalStakedUsdc);
    }

    function viewStaker() public view returns (uint256, uint256, uint256, bool) {
        Staker storage staker = stakers[msg.sender];
        return (staker.ethStaked, staker.usdtStaked, staker.usdcStaked, staker.isMintDice);
    }

    function getDecimals(address token) public view returns (uint8) {
        ERC20 erc20 = ERC20(token);
        return erc20.decimals();
    }
}