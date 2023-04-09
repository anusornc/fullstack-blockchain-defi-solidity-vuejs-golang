// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

// LPToken is a token that represents the liquidity pool
contract LPToken is Ownable, ERC20 {
  constructor() ERC20('LPToken', 'LPToken') {
  }
}