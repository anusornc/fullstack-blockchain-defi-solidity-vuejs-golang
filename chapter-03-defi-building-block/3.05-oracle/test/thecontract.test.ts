import expect from 'expect'
import ganache from 'ganache-cli'
import { AbiItem } from 'web3-utils'
import { TheOracle } from '../types/web3-v1-contracts/TheOracle'
import TheOracleContract from '../build/contracts/TheOracle.json'
import { TheContract } from '../types/web3-v1-contracts/TheContract'
import TheContractContract from '../build/contracts/TheContract.json'
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

let theOracle: TheOracle
let theContract: TheContract

beforeEach(async () => {
  const actors = await testActors(web3)   // Get test accounts

  // 1. Deploy Oracle contract
  theOracle = new web3.eth.Contract(TheOracleContract.abi as AbiItem[]) as any as TheOracle // Cast to TheOracle
  theOracle = await theOracle.deploy({data: TheOracleContract.bytecode, arguments:[]})  // Deploy
    .send(actors.ownerTx) as any as TheOracle

  // 2. Get oracle address
  const theOracleAddr = theOracle.options.address // Get oracle address

  // 3. Deploy theContract contract
  theContract = new web3.eth.Contract(TheContractContract.abi as AbiItem[]) as any as TheContract   // Cast to TheContract
  theContract = await theContract.deploy({data: TheContractContract.bytecode, arguments:[theOracleAddr]}) // Deploy
    .send(actors.ownerTx) as any as TheContract
})

describe('The Contract', () => {
  it('can use price from oracle', async () => {
    const actors = await testActors(web3) // Get test accounts
    const key = toBytes32('BTC/DAI')    // Get key for BTC/DAI
    await theOracle.methods.addUpdater(actors.acc1Addr, true).send(actors.ownerTx)  // Add acc1 as updater
    // 1 BTC = 50,000 DAI
    await theOracle.methods.updatePrice(key, '50000').send(actors.acc1Tx) // Update price

    // Try swap BTC/DAI from theContract 
    const poolId = toBytes32('BTC/DAI') // Get poolId for BTC/DAI
    await theContract.methods.swap(poolId, toWei('1')).send(actors.acc2Tx)  // Swap 1 BTC for DAI

    // Validate the result from events (swap 1 BTC for 50K DAI)
    // Q: What is events do ?
    // A: events is an array of objects that contains the information of the event
    // Q: What is theContract.getPastEvents('Swap', {filter: {}, fromBlock: 0}) do ?
    // A: It returns an array of objects that contains the information of the event
    // Q: Can we rename events to something else ?
    // A: Yes, we can rename it to anything we want
    const events = await theContract.getPastEvents('Swap', {
      filter: {},
      fromBlock: 0,
    })
    expect(events.length).toEqual(1)    // Check event count
    expect(events[0].returnValues['sender']).toEqual(actors.acc2Addr) // Check sender
    expect(fromWei(events[0].returnValues['amountIn'])).toEqual('1')  // Check amountIn
    expect(fromWei(events[0].returnValues['amountOut'])).toEqual('50000') // Check amountOut

    // update price to 100,000 DAI
    await theOracle.methods.updatePrice(key, '100000').send(actors.acc1Tx) // Update price

    // Try swap BTC/DAI from theContract
    await theContract.methods.swap(poolId, toWei('1')).send(actors.acc2Tx)  // Swap 1 BTC for DAI

    // Validate the result from events (swap 1 BTC for 100K DAI)
    const events2 = await theContract.getPastEvents('Swap', {
      filter: {},
      fromBlock: 0,
    })

    console.log(events2.length)

    expect(events2.length).toEqual(2)   // Check event count
    expect(events2[1].returnValues['sender']).toEqual(actors.acc2Addr) // Check sender
    expect(fromWei(events2[0].returnValues['amountIn'])).toEqual('1')  // Check amountIn
    expect(fromWei(events2[0].returnValues['amountOut'])).toEqual('50000') // Check amountOut
    
    expect(fromWei(events2[1].returnValues['amountIn'])).toEqual('1')  // Check amountIn
    expect(fromWei(events2[1].returnValues['amountOut'])).toEqual('100000') // Check amountOut


  })
})