import expect from 'expect'
import ganache from 'ganache-cli'
import { AbiItem } from 'web3-utils'
import { BUSD } from '../types/web3-v1-contracts/BUSD'
import BUSDContract from '../build/contracts/BUSD.json'
import { MyToken } from '../types/web3-v1-contracts/MyToken'
import MyTokenContract from '../build/contracts/MyToken.json'
import {beforeEach, describe, it, jest} from '@jest/globals'
import {
  getWeb3, 
  toWei, 
  toBasis, 
  toBN, 
  fromWei,
  zeroBN,
  toBytes32,
  fromBytes32,
} from '../utils/util'
import { 
  testActors
} from '../utils/test_util'

const ganacheOpts = {
  // verbose: true,
  // logger: console,
}
const web3 = getWeb3(ganache.provider(ganacheOpts))

let assetToken: BUSD
let myToken: MyToken

beforeEach(async () => {
  const actors = await testActors(web3) // Get test actors
  // 1. Deploy BUSD faucet as asset token
  assetToken = new web3.eth.Contract(BUSDContract.abi as AbiItem[]) as any as BUSD // Create contract instance
  assetToken = await assetToken.deploy({data: BUSDContract.bytecode, arguments:[]}) // Deploy contract
    .send(actors.ownerTx) as any as BUSD

  // 2. Get address of back asset
  const assetAddr = assetToken.options.address // Get address of asset token
  
  // 3. Deploy myToken contract
  myToken = new web3.eth.Contract(MyTokenContract.abi as AbiItem[]) as any as MyToken // Create contract instance
  myToken = await myToken.deploy({data: MyTokenContract.bytecode, arguments:[assetAddr]}) // Deploy contract
    .send(actors.ownerTx) as any as MyToken
})

describe('My Token', () => {
  it('can mint MyToken backed by asset', async () => {
    const actors = await testActors(web3) // Get test actors

    const myTokenAddr = myToken.options.address // Get address of MyToken contract

    // 1. Mint 100 BUSD for test
    const amount = toWei('100')   // 100 BUSD
    await assetToken.methods.mint(amount).send(actors.acc1Tx) // Mint 100 BUSD for acc1

    // Validate that acc1 will have 100 BUSD and 0 MyToken
    let assetBalAcc1 = await assetToken.methods.balanceOf(actors.acc1Addr).call() // acc1's BUSD balance
    let assetBalContract = await assetToken.methods.balanceOf(myTokenAddr).call() // MyToken contract's BUSD balance
    let myTokenBalAcc1 = await myToken.methods.balanceOf(actors.acc1Addr).call()  // acc1's MyToken balance

    expect(fromWei(assetBalAcc1)).toEqual('100')  // acc1's BUSD balance
    expect(fromWei(assetBalContract)).toEqual('0') // MyToken contract's BUSD balance
    expect(fromWei(myTokenBalAcc1)).toEqual('0') // acc1's MyToken balance

    // Make deposit (exchange 100 BUSD for 100 MyToken)
    await assetToken.methods.approve(myTokenAddr, amount).send(actors.acc1Tx) // Approve MyToken contract to spend 100 BUSD
    await myToken.methods.deposit(amount).send(actors.acc1Tx) // Deposit 100 BUSD

    // Validate that acc1 will have 0 BUSD and 100 MyToken
    // and MyToken contract will have 100 BUSD back the 100 MyToken
    assetBalAcc1 = await assetToken.methods.balanceOf(actors.acc1Addr).call()
    assetBalContract = await assetToken.methods.balanceOf(myTokenAddr).call()
    myTokenBalAcc1 = await myToken.methods.balanceOf(actors.acc1Addr).call()

    expect(fromWei(assetBalAcc1)).toEqual('0')
    expect(fromWei(assetBalContract)).toEqual('100')
    expect(fromWei(myTokenBalAcc1)).toEqual('100')


    await myToken.methods.withdraw(amount).send(actors.acc1Tx)
    assetBalAcc1 = await assetToken.methods.balanceOf(actors.acc1Addr).call()
    assetBalContract = await assetToken.methods.balanceOf(myTokenAddr).call()
    myTokenBalAcc1 = await myToken.methods.balanceOf(actors.acc1Addr).call()

    expect(fromWei(assetBalAcc1)).toEqual('100')
    expect(fromWei(assetBalContract)).toEqual('0')
    expect(fromWei(myTokenBalAcc1)).toEqual('0')

  })


})