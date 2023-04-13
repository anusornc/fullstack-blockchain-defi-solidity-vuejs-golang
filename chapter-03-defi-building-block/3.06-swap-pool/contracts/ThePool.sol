// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import './LPToken.sol';
import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import "@openzeppelin/contracts/utils/math/Math.sol";

<<<<<<< HEAD
// ThePool is a Uniswap-like swap pool
contract ThePool is LPToken {
  IERC20 private _tokenA;   // tokenA is the base token
  IERC20 private _tokenB;   // tokenB is the quote token
=======

// Q: What is the Swap Pool?
// A: The Swap Pool is a smart contract that holds the assets of the liquidity provider.
//    The liquidity provider can deposit the assets into the swap pool and get the liquidity pool token.
//    The liquidity provider can withdraw the assets from the swap pool by burning the liquidity pool token.
//    The liquidity provider can swap the assets in the swap pool.
//    The swap pool is owned by the liquidity provider.
//    The swap pool is minted by the liquidity provider.
//    The swap pool is burned by the liquidity provider.
//    The swap pool is swapped by the liquidity provider.

contract ThePool is LPToken {
  IERC20 private _tokenA; 
  IERC20 private _tokenB;
>>>>>>> 2a6043d (update util and other)

  constructor(address tokenA, address tokenB) LPToken() {
    _tokenA = IERC20(tokenA);
    _tokenB = IERC20(tokenB);
  }

  // add liquidity to the pool by depositing tokenA and tokenB into the pool
  // the liquidity provider gets the liquidity pool token
  // the liquidity pool token is minted by the liquidity pool
  // the liquidity pool token is owned by the liquidity provider
  // the liquidity pool token is burned by the liquidity provider
  function addLiquidity(uint amountA, uint amountB) public {

    address contractAddr = address(this); // the address of the liquidity pool
    uint reserveA = _tokenA.balanceOf(contractAddr); // the balance of tokenA in the liquidity pool 
    uint reserveB = _tokenB.balanceOf(contractAddr); // the balance of tokenB in the liquidity pool
    
    // transfer tokenA and tokenB from the liquidity provider to the liquidity pool
    _tokenA.transferFrom(msg.sender, contractAddr, amountA); 

    // transfer tokenA and tokenB from the liquidity provider to the liquidity pool
    _tokenB.transferFrom(msg.sender, contractAddr, amountB);

    uint newReserveA = _tokenA.balanceOf(contractAddr); // the balance of tokenA in the liquidity pool
    uint newReserveB = _tokenB.balanceOf(contractAddr); // the balance of tokenB in the liquidity pool

    // reserveA * reserveB = k
    // newReserveA * newReserveB >= k
    // newReserveA * newReserveB >= reserveA * reserveB
    require(newReserveA * newReserveB >= reserveA * reserveB, 'invalid K'); 

    // the total supply of the liquidity pool token
    uint totalSupply = totalSupply();

    // the amount of liquidity pool token to be minted
    uint lpTokenAmount = 0;

    // if the total supply of the liquidity pool token is 0
    // the amount of liquidity pool token to be minted is the amount of tokenA
    // else
    // the amount of liquidity pool token to be minted is the minimum of
    // the amount of tokenA * the total supply of the liquidity pool token / the balance of tokenA in the liquidity pool
    // and the amount of tokenB * the total supply of the liquidity pool token / the balance of tokenB in the liquidity pool

    if (totalSupply == 0) {
      lpTokenAmount = amountA;
    } else {
      lpTokenAmount = Math.min(amountA * totalSupply / reserveA, 
                                amountB * totalSupply / reserveB);
    }
    
    // the amount of liquidity pool token to be minted must be greater than 0
    require(lpTokenAmount > 0, 'lptoken amount == 0');

    // mint the liquidity pool token to the liquidity provider
    _mint(msg.sender, lpTokenAmount);
  }

  // remove liquidity from the pool by burning the liquidity pool token
  // the liquidity provider gets the tokenA and tokenB from the pool
  // the liquidity pool token is burned by the liquidity pool
  // the liquidity pool token is owned by the liquidity provider
  // the liquidity pool token is minted by the liquidity provider

  function removeLiquidity(uint amount) public {
    // the liquidity provider must have the liquidity pool token
    require(balanceOf(msg.sender) >= amount, 'insufficient balance');

    
    address contractAddr = address(this); // the address of the liquidity pool

    uint reserveA = _tokenA.balanceOf(contractAddr); // the balance of tokenA in the liquidity pool
    uint reserveB = _tokenB.balanceOf(contractAddr); // the balance of tokenB in the liquidity pool

    // the amount of tokenA to be transferred to the liquidity provider
    // the amount of tokenA to be transferred to the liquidity provider is the amount of liquidity pool token * the balance of tokenA in the liquidity pool / the total supply of the liquidity pool token
    
    uint totalSupply = totalSupply(); // the total supply of the liquidity pool token
    uint amountA = amount * reserveA / totalSupply; 
    uint amountB = amount * reserveB / totalSupply;

    _burn(msg.sender, amount); // burn the liquidity pool token from the liquidity provider
    _tokenA.transfer(msg.sender, amountA); // transfer tokenA to the liquidity provider
    _tokenB.transfer(msg.sender, amountB); // transfer tokenB to the liquidity provider
  }

  // swap tokenA for tokenB
  // the liquidity provider gets the tokenB from the pool
  // the liquidity provider gives the tokenA to the pool
  // the liquidity pool token is burned by the liquidity pool
  // the liquidity pool token is owned by the liquidity provider
  // the liquidity pool token is minted by the liquidity provider

  function swapAForB(uint amountAIn, uint amountBOut) public {

    require(amountAIn > 0, 'insufficient amountA'); // the amount of tokenA to be swapped must be greater than 0
    require(amountBOut > 0, 'insufficient amountB'); // the amount of tokenB to be swapped must be greater than 0

    address contractAddr = address(this); // the address of the liquidity pool

    uint reserveA = _tokenA.balanceOf(contractAddr); // the balance of tokenA in the liquidity pool
    uint reserveB = _tokenB.balanceOf(contractAddr); // the balance of tokenB in the liquidity pool
    
    // the amount of tokenA to be swapped must be less than or equal to the balance of tokenA in the liquidity pool
    require(amountAIn <= reserveA, 'insufficient reserveA'); // [** added ] the amount of tokenA to be swapped must be less than or equal to the balance of tokenA in the liquidity pool
    _tokenA.transferFrom(msg.sender, address(this), amountAIn); // transfer tokenA from the liquidity provider to the liquidity pool
    _tokenB.transfer(msg.sender, amountBOut); // transfer tokenB from the liquidity pool to the liquidity provider

    uint newReserveA = _tokenA.balanceOf(contractAddr); // the balance of tokenA in the liquidity pool 
    uint newReserveB = _tokenB.balanceOf(contractAddr); // the balance of tokenB in the liquidity pool

    // reserveA * reserveB = k
    // newReserveA * newReserveB >= k
    require(newReserveA * newReserveB >= reserveA * reserveB, 'invalid K'); // 
  }

  // swap tokenB for tokenA
  // the liquidity provider gets the tokenA from the pool
  // the liquidity provider gives the tokenB to the pool
  // the liquidity pool token is burned by the liquidity pool
  // the liquidity pool token is owned by the liquidity provider
  // the liquidity pool token is minted by the liquidity provider

  function getAmountBByA(uint amountA) public view returns (uint) {
    require(amountA > 0, 'insufficient amount');

    address contractAddr = address(this); // the address of the liquidity pool
    uint reserveA = _tokenA.balanceOf(contractAddr); // the balance of tokenA in the liquidity pool
    uint reserveB = _tokenB.balanceOf(contractAddr); // the balance of tokenB in the liquidity pool
    require(reserveA > 0 && reserveB > 0, 'insufficient liquidity'); // the balance of tokenA and tokenB in the liquidity pool must be greater than 0

    return amountA * reserveB / reserveA;
  }

  function swapBForA(uint amountBIn, uint amountAOut) public {

    require(amountBIn > 0, 'insufficient amountB');
    require(amountAOut > 0, 'insufficient amountA');

    address contractAddr = address(this);

    uint reserveA = _tokenA.balanceOf(contractAddr);
    uint reserveB = _tokenB.balanceOf(contractAddr);
    
    _tokenA.transfer(msg.sender, amountAOut);
    _tokenB.transferFrom(msg.sender, address(this), amountBIn);

    uint newReserveA = _tokenA.balanceOf(contractAddr);
    uint newReserveB = _tokenB.balanceOf(contractAddr);

    // reserveA * reserveB = k
    // newReserveA * newReserveB >= k
    require(newReserveA * newReserveB >= reserveA * reserveB, 'invalid K');
  }

  function getAmountAByB(uint amountB) public view returns (uint) {
    require(amountB > 0, 'insufficient amount');

    address contractAddr = address(this);
    uint reserveA = _tokenA.balanceOf(contractAddr);
    uint reserveB = _tokenB.balanceOf(contractAddr);
    require(reserveA > 0 && reserveB > 0, 'insufficient liquidity');

    return amountB * reserveA / reserveB;
  }
}