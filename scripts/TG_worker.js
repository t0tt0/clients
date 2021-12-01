const HELP_MSG =
  'Usage: node TG_worker.js -t [T] -c [N]\n\n\t-c [N]\tDeploy a contract to test.\n\t\tN\tNumber of transactions generated to the contract per second.\n\n\t-t [T]\tSet test duration.\n\t\tT\tTest duration in seconds.\n\n\t -i [I]\tWorker id.\n\t\tI\tId of the worker to be shown in the logs.\n\n\t-h\tDisplay this message.'

const TEST_CONTRACT_1 =
  'contract testContract1{ ' +
  '    struct CallRecord{ ' +
  '        uint callid; ' +
  '        address from; ' +
  '        uint time; ' +
  '    } ' +
  '    CallRecord[] public callRecords; ' +
  '    uint public numberOfCallRecords; ' +
  '    address owner; ' +
  '    constructor() public{ ' +
  '        owner = msg.sender; ' +
  '        numberOfCallRecords = 0; ' +
  '    } ' +
  '    function injectTx(uint callid) public{ ' +
  '        callRecords.push(CallRecord({ ' +
  '            callid: callid, ' +
  '            from: msg.sender, ' +
  '            time: block.timestamp ' +
  '        })); ' +
  '    } ' +
  '} '

const accounts = [
  {
    address: '0xd3d091e01502d53e3820aed1e6eff3af2673f346',
    privateKey: '0xc17841232ec7c6654ffd79ae3a8bbeec9da63b99fbdf42dffdb820f4d734d0ca'
  },
  {
    address: '0x94abee207bc81159831e3c6b3a6722d019e8e13e',
    privateKey: '0xa24e4d4657fb5870031d602d0231f14f877ffd8e132af70cd72a938201784466'
  },
  {
    address: '0xe85043677AC0dfe83277Cd83E52CE2d43e871Fa1',
    privateKey: '0x95c0e15b8bec8ed0365d2ccbfb25de3b4c233ae593df0338f1ef85b79c0eb500'
  },
  {
    address: '0x05d739c60C913689634429A393988d88B685183F',
    privateKey: '0xbafa81e271c5db61ca28a1f6d3885b66e2fbcf71872947bdfbc950a0b9dd348e'
  },
  {
    address: '0xF1075F6AF05F7043782D3934204175F57E0cF50D',
    privateKey: '0xaf602c78b7a3dc9103e64a6bf7bb0ae48657fa834221c5da19184fc962de4141'
  },
  {
    address: '0xD8D01FA5aDc59cBFacB5eC876a5c5AB15DfB4e3f',
    privateKey: '0x37a538c2db9a8dd2b96ade0b7ab652b420106180281a1fcf166287aa4995b81d'
  },
  {
    address: '0x99Aa2732bFf87fD7797861734e58B6e8C5dA43be',
    privateKey: '0xaedb414cd36d9bb0b2db3deb95218f195a2dfca09a9fc53e041b590ed5d5393b'
  },
  {
    address: '0x18130B61a85c83d4355C67Bd31e5Dc6fd8e300E8',
    privateKey: '0xf66a1ffb5e172e5e1d4ce63ac42d804e2ce6555645ccc598cf5ff4090c01dadf'
  }
]

var Web3 = require('web3')
var solc = require('solc')
var fs = require('fs')
var net = require('net')

function transactionInjector (
  web3,
  contract,
  account,
  log,
  workerId,
  duration,
  pace,
  nonce,
  counter = 0
) {
  var delay = 1000 / pace
  log.consoleLog(workerId, 'Transaction ' + counter + ' injecting.')
  contract.methods
    .injectTx(counter)
    .send({ from: account.address, gas: 1000000, 
nonce: nonce, gasprice: (counter + 1)*2*1000000000 })
    .on('receipt', recp =>
      log.consoleLog(workerId, 'Transaction ' + counter + ' injected.')
    )
  if (counter < duration * pace) {
    setTimeout(() => {
      transactionInjector(
        web3,
        contract,
        account,
        log,
        workerId,
        duration,
        pace,
        nonce + 1,
        counter + 1
      )
    }, delay)
  } else {
  }
}

function worker (argv) {
  var args = require('minimist')(argv.slice(2))

  if (
    args['help'] ||
    args['h'] ||
    args['t'] === undefined ||
    isNaN(args['t']) ||
    args['c'] === undefined ||
    isNaN(args['c']) ||
    args['i'] === undefined ||
    isNaN(args['i'])
  ) {
    console.log(HELP_MSG)
    return
  }

  var testDuration = args['t']
  var contractLoad = args['c']
  var workerId = args['i']

  var log = require('./TG_log')

  log.consoleLog(
    workerId,
    'Started with ' +
      contractLoad +
      ' tx/sec load for ' +
      testDuration +
      ' seconds.'
  )

  var web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:24765'))
  log.consoleLog(workerId, 'Connected to Ethereum node.')
  web3.eth
    .getBlockNumber()
    .then(blkNum => log.consoleLog(workerId, 'Block number: ' + blkNum + '.'))
    .catch(() => log.consoleLog(workerId, 'Ethereum node does not reply.'))

  var account = web3.eth.accounts.privateKeyToAccount(
    accounts[parseInt(workerId)].privateKey
  )
  web3.eth.accounts.wallet.add(account)

  log.consoleLog(workerId, 'Ethereum account ' + account.address + ' used.')
  log.consoleLog(workerId, 'Ethereum account private key ' + account.privateKey)

  var testContract = {
    'testContract1.sol': TEST_CONTRACT_1
  }

  var solcOut = solc.compile({ sources: testContract }, 1)

  var compiledContract = {}

  for (var contractName in solcOut.contracts) {
    compiledContract = solcOut.contracts[contractName]
  }
  compiledContract.parsedInterface = JSON.parse(compiledContract.interface)

  var ethContract = new web3.eth.Contract(compiledContract.parsedInterface)

  var ethContractDeployment = ethContract
    .deploy({ data: '0x' + compiledContract.bytecode })
    .send({ from: account.address, gas: 1000000 })
    .then(contract => {
      log.consoleLog(
        workerId,
        'Deployed contract address: ' + contract.options.address
      )
      web3.eth.getTransactionCount(account.address).then(nonce => {
        log.consoleLog(workerId, nonce)
        transactionInjector(
          web3,
          contract,
          account,
          log,
          workerId,
          testDuration,
          contractLoad,
          nonce,
          0
        )
      })
    })
  // .catch(() =>
  //     log.consoleLog(workerId, "Contract deployment failed."));
}

worker(process.argv)
