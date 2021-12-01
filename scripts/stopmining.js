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
    address: '0xd3d091e01502d53e3820aed1e6eff3af2673f346', //eth1
    privateKey: '0xc17841232ec7c6654ffd79ae3a8bbeec9da63b99fbdf42dffdb820f4d734d0ca'
  },
  {
    address: '0x94abee207bc81159831e3c6b3a6722d019e8e13e', //eth1[1]
    privateKey: '0xa24e4d4657fb5870031d602d0231f14f877ffd8e132af70cd72a938201784466'
  },
  {
    address: '0xb782edeacc48fa6161dc690043bd0db17e959bf4', //ves1[0]
    privateKey: '0x93c4bb70eafb393e6b7d338b0e889de33925935d77fd3fe3540a99c8203378a8'
  },
  {
    address: '0xea2ba9cfae53f04989faefeb6fa7f0b65a72e4b0', //ves1[1]
    privateKey: '0x76e3237d64c57ef50da5d78ff3024942962a439b86ab4cf50bf2c38339ceb1ac'
  },
  {
    address: '0x2d05a3fbe28aadd143701727c8a766694654fb99', //ves2[0]
    privateKey: '0xf2add373ac0979002383c4a41db3bda1278492b4a6764f1a53dc43a73dee4b69'
  },
  {
    address: '0x6f55fcf81624e908da2ba3c83341ff707b652760', //ves2[1]
    privateKey: '0xe45d0d1fc072ed0272cda33c9c2021beb4e4db47e01d471227db6b4ee54cb627'
  },
  {
    address: '0xc157a9260be13546d83bd089450e2cc25b38ae29', //eth2[0]
    privateKey: '0x5625e3268511e22e021a03a713132b0cfcaf47571a8e31a7328129f81a1f5424'
  },
  {
    address: '0xe790a5c30ccaef61bc123326bd3fc65a7f039ba7', //eth2[1]
    privateKey: '5625e3268511e22e021a03a713132b0cfcaf47571a8e31a7328129f81a1f5424'
  }
]

var Web3 = require('web3')
var solc = require('solc')
var fs = require('fs')
var net = require('net')
var path = require('path')


function anotherworker () {
  workerId = 5

  var log = require('./TG_log')

  var web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:24768'))
  log.consoleLog(workerId, 'Connected to another Ethereum node.')
  web3.eth
    .getBlockNumber()
    .then(blkNum => log.consoleLog(workerId, 'Block number: ' + blkNum + '.'))
    .catch(() => log.consoleLog(workerId, 'Ethereum node does not reply.'))


  // var keythereum = require("keythereum");
  // var datadir = "/Users/taoyuechen/Library/Ethereum1";
  // var address= "0x94abee207bc81159831e3c6b3a6722d019e8e13e";
  // const password = "1";

  // var keyObject = keythereum.importFromFile(address, datadir);
  // var privateKey = keythereum.recover(password, keyObject);
  // log.consoleLog(workerId, privateKey.toString('hex'));

  var account1 = web3.eth.accounts.privateKeyToAccount(
    accounts[parseInt(workerId)].privateKey
  )
  web3.eth.accounts.wallet.add(account1);

  var account2 = web3.eth.accounts.privateKeyToAccount(
    accounts[parseInt(4)].privateKey
  )

  log.consoleLog(workerId, 'Ethereum account ' + account1.address + ' used.')
  log.consoleLog(workerId, 'Ethereum account private key ' + account1.privateKey)

  const rlp = require('rlp');
  const keccak = require('keccak');

  var nonce = 0x00; //The nonce must be a hex literal!
  var sender = account2.address; //Requires a hex string as input!

  var input_arr = [ sender, nonce ];
  var rlp_encoded = rlp.encode(input_arr);

  var contract_address_long = keccak('keccak256').update(rlp_encoded).digest('hex');

  var contract_address = contract_address_long.substring(24);
  log.consoleLog(workerId, "sender in chain2 "+sender)
  log.consoleLog(workerId, "contract_address in chain 2 "+contract_address)

  web3.eth.getBalance(contract_address)
  .then(blc => {
    if (blc > 0) {
      log.consoleLog(workerId, 'y.....'+blc)
      web3.eth.sendTransaction({
        from: account1.address,
        to: contract_address,
        gas: 215650,
        data: "0xf1768b6c0000000000000000000000000000000000000000000000000000000000000001"
      })
    }
    else {
      log.consoleLog(workerId, 'n.....'+blc)
      web3.eth.sendTransaction({
        from: account1.address,
        to: contract_address,
        gas: 215650,
        data: "0xf1768b6c0000000000000000000000000000000000000000000000000000000000000002"
      })
    }
  });

  // web3.eth.getTransactionCount(accounts[2].address).then(after => 
  //   if (after >= 2) {
  //     log.consoleLog(workerId, after)
  //   }
  //   //todo 
  //   web3.eth.sendTransaction({
  //     from: account.address,
  //     gas: 1000000,
  //     // value: 60000000000000,
  //    // data:  // invoking functions, further conducted
  //   }, function(error, hash){
  //     }
  // )

  
}


function worker () {

  // if (args['c'] !== undefined && !Array.isArray(args['c'])) {
  //   args['c'] = [args['c']]
  // }

  // if (
  //   args['help'] ||
  //   args['h'] ||
  //   args['t'] === undefined ||
  //   isNaN(args['t']) ||
  //   args['c'] === undefined ||
  //   !args['c'].every(i => !isNaN(i))
  // ) {
  //   console.log(HELP_MSG)
  //   return
  // }


  
  workerId = 0

  var log = require('./TG_log')

  var web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:24765'))
  var web3admin = require('./web3Admin.js')
  log.consoleLog(workerId, 'Connected to Ethereum node.')

 
  setTimeout(function(){
    web3admin.extend(web3)
 
    log.consoleLog(workerId, 'mining?     '+web3.eth.mining)
    web3.miner.stop()
    log.consoleLog(workerId, 'mining?     '+web3.eth.mining)
  }, 1)

  
  var web31 = new Web3(new Web3.providers.HttpProvider('http://localhost:24768'))
  var web3admin1 = require('./web3Admin.js')
  log.consoleLog(workerId, 'Connected to Another Ethereum node.')

 
  setTimeout(function(){
    web3admin1.extend(web31)
 
    log.consoleLog(workerId+1, 'mining?     '+web31.eth.mining)
    web31.miner.stop()
    log.consoleLog(workerId, 'mining?     '+web31.eth.mining)
  }, 1)

  

 

}


//todo: hyperservice: web3, change to launching smart contracts..., write a script for gettointial, and send transactions to hyperservice; write a contract on hyperservice
function main (argv) {
  // args = require('minimist')(argv.slice(2))
  // if (args['c'] !== undefined && !Array.isArray(args['c'])) {
  //   args['c'] = [args['c']]
  // }

  // if (
  //   args['help'] ||
  //   args['h'] ||
  //   args['t'] === undefined ||
  //   isNaN(args['t']) ||
  //   args['c'] === undefined ||
  //   !args['c'].every(i => !isNaN(i))
  // ) {
  //   console.log(HELP_MSG)
  //   return
  // }

  // // var log = require('./TG_log')
  // // const rlp = require('rlp');
  // // const keccak = require('keccak');

  // var nonce = args['t']

  // account1 = accounts[parseInt(0)].address

  // var sender = account1; //Requires a hex string as input!
  // log.consoleLog(0, 'nonce is ' + nonce)

  // var input_arr = [ sender, nonce ];
  // var rlp_encoded = rlp.encode(input_arr);

  // var contract_address_long = keccak('keccak256').update(rlp_encoded).digest('hex');

  // var contract_address = contract_address_long.substring(24);

  // log.consoleLog(0, 'contract address is ' + contract_address)

  


  worker()
}

main(process.argv)
// worker(process.argv)
