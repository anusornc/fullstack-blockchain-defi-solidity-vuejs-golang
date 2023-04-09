// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import './IOracle.sol';

contract TheContract {
  
  IOracle private _oracle;

  // Q: What is the purpose of the event?
  // A: The event is used to log the swap function's input and output.
  //   The event is emitted when the swap function is called.
  
  event Swap(address indexed sender, uint amountIn, uint amountOut);  // event for swap function

  constructor(address oracleAddr) {
    _oracle = IOracle(oracleAddr);
  }

  // Q: What is the purpose of the swap function?
  // A: The swap function is used to swap tokens.
  //   The swap function is called when a user wants to swap tokens.
  //   The swap function takes the poolId and amountIn as input.
  //   The swap function calls the getPrice function of the oracle contract to get the price of the pool.
  //   The swap function checks if the price is outdated.
  //   The swap function checks if the price is set.
  //   The swap function calculates the amountOut.
  //   The swap function emits the Swap event.

  function swap(bytes32 poolId, uint amountIn) public {
    IOracle.Price memory price = _oracle.getPrice(poolId);    // get price from oracle
    require(block.number - price.blocknumber < 10, 'price is outdated'); // check if price is outdated
    require(price.data > 0, 'no price set'); // check if price is set
    uint amountOut = amountIn * price.data; // calculate amountOut
    
    // Q: What is the purpose of the emit statement?
    // A: The emit statement is used to emit the Swap event.
    //   The emit statement is called when the swap function is called.
    //   The emit statement emits the Swap event with the input and output values.
    
    emit Swap(msg.sender, amountIn, amountOut); // emit event for swap function call with input and output values 
  }
}