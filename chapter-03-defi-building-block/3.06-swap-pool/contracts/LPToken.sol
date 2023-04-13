// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

<<<<<<< HEAD
// LPToken is a token that represents the liquidity pool
=======
// Q: What is the Liquidity Pool Token?
// A: The Liquidity Pool Token is a token that represents the share of the liquidity pool.
//    The liquidity pool is a smart contract that holds the assets of the liquidity provider.
//    The liquidity provider can deposit the assets into the liquidity pool and get the liquidity pool token.
//    The liquidity provider can withdraw the assets from the liquidity pool by burning the liquidity pool token.
//    The liquidity pool token is an ERC20 token.
//    The liquidity pool token is minted when the liquidity provider deposits the assets into the liquidity pool.
//    The liquidity pool token is burned when the liquidity provider withdraws the assets from the liquidity pool.
//    The liquidity pool token is owned by the liquidity pool.
//    The liquidity pool token is minted by the liquidity pool.
//    The liquidity pool token is burned by the liquidity pool.

>>>>>>> 2a6043d (update util and other)
contract LPToken is Ownable, ERC20 {
  constructor() ERC20('LPToken', 'LPToken') {
  }
}