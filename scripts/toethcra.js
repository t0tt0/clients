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


function worker (nonce) {

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


  
  workerId = 1

  var log = require('./TG_log')

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

  // var account1 = web3.eth.accounts.privateKeyToAccount(
  //   accounts[parseInt(0)].privateKey
  // )

  log.consoleLog(workerId, 'Ethereum account ' + account.address + ' used.')
  log.consoleLog(workerId, 'Ethereum account private key ' + account.privateKey)









  var web31 = new Web3(new Web3.providers.HttpProvider('http://localhost:24768'))
  log.consoleLog(workerId, 'Connected to another Ethereum node.')
  web31.eth
    .getBlockNumber()
    .then(blkNum => log.consoleLog(workerId, 'Block number: ' + blkNum + '.'))
    .catch(() => log.consoleLog(workerId, 'Ethereum node does not reply.'))

  var account2 = web31.eth.accounts.privateKeyToAccount(
    accounts[parseInt(5)].privateKey
  )
  web31.eth.accounts.wallet.add(account2)

  // var account3 = web31.eth.accounts.privateKeyToAccount(
  //   accounts[parseInt(4)].privateKey
  // )

  log.consoleLog(workerId, 'another Ethereum account ' + account2.address + ' used.')
  log.consoleLog(workerId, 'another Ethereum account private key ' + account2.privateKey)



  const rlp = require('rlp');
  const keccak = require('keccak');

  // var nonce = 0x00; //The nonce must be a hex literal!
  var sender = account.address; //Requires a hex string as input!

  var input_arr = [ sender, nonce ];
  var rlp_encoded = rlp.encode(input_arr);

  var contract_address_long = keccak('keccak256').update(rlp_encoded).digest('hex');

  var contract_address = contract_address_long.substring(24);


  //var nonce = 0x00; //The nonce must be a hex literal!
  var sender1 = account2.address; //Requires a hex string as input!

  var input_arr1 = [ sender1, nonce ];
  var rlp_encoded1 = rlp.encode(input_arr1);

  var contract_address_long1 = keccak('keccak256').update(rlp_encoded1).digest('hex');

  var contract_address1 = contract_address_long1.substring(24);

  log.consoleLog(workerId, 'contract_address is ' + contract_address)
  log.consoleLog(workerId, 'contract_address1 is ' + contract_address1)

  var bytecode = '608060405233600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600080819055506102fe8061005b6000396000f3fe6080604052600436106100345760003560e01c8063202ec1521461003957806382ab890a14610043578063da7d29821461006c575b600080fd5b610041610088565b005b34801561004f57600080fd5b5061006a6004803603810190610065919061022c565b61010f565b005b610086600480360381019061008191906101ff565b61016f565b005b6002600054141561010d57600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc670de0b6b3a76400009081150290604051600060405180830381858888f19350505050158015610103573d6000803e3d6000fd5b5060066000819055505b565b6005600054148061012257506006600054145b8061012f57506001600054145b8061013c57506002600054145b1561014f5760005460008190555061016c565b600181148061015e5750600281145b1561016b57806000819055505b5b50565b600160005414156101d2578073ffffffffffffffffffffffffffffffffffffffff166108fc670de0b6b3a76400009081150290604051600060405180830381858888f193505050501580156101c8573d6000803e3d6000fd5b5060056000819055505b50565b6000813590506101e48161029a565b92915050565b6000813590506101f9816102b1565b92915050565b60006020828403121561021557610214610295565b5b6000610223848285016101d5565b91505092915050565b60006020828403121561024257610241610295565b5b6000610250848285016101ea565b91505092915050565b60006102648261026b565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600080fd5b6102a381610259565b81146102ae57600080fd5b50565b6102ba8161028b565b81146102c557600080fd5b5056fea2646970667358221220dec7558ee50777622ebf86d566958cf3fa5d2016d23cf4949afebe49349bdc1464736f6c63430008070033';


  web3.eth.sendTransaction({
    from: account.address,
    gas: 398928,
    value: 3000000000000000000,
    data: '0x'+ bytecode
  }, function(error, hash) {
//to be added
  }).on('transactionHash', function(error, hash){

  });

  

  web31.eth.sendTransaction({
    from: account2.address,
    gas: 398928,
    value: 3000000000000000000,
    data: '0x'+ bytecode
  }, function(error, hash) {
//to be added
  }).on('transactionHash', function(error, hash){

  });

  // web3.eth.sendTransaction({from: account.address, to: contract_address, value: 200000, gas: 23000})
  // web31.eth.sendTransaction({from: account2.address, to: contract_address1, value: 200000, gas: 23000})

  // web3.eth.getBalance(contract_address)
  // .then(blc => {
  //   if(blc > 0) {
  //     web31.eth.getBalance(contract_address1)
  //     .then(blc1 => {
  //       if(blc1 > 0) {
  //         log.consoleLog(workerId, "y....."+blc + " "+blc1)
  //         web3.eth.sendTransaction({
  //           from: account.address,
  //           to: contract_address,
  //           gas: 215650,
  //           data: "0xf1768b6c0000000000000000000000000000000000000000000000000000000000000001"
  //         });
  //         web31.eth.sendTransaction({
  //           from: account2.address,
  //           to: contract_address1,
  //           gas: 215650,
  //           data: "0xf1768b6c0000000000000000000000000000000000000000000000000000000000000001"
  //         });
  //       }
  //       else {
  //         log.consoleLog(workerId, "n....."+blc + " "+blc1)
  //         web3.eth.sendTransaction({
  //           from: account.address,
  //           to: contract_address,
  //           gas: 215650,
  //           data: "0xf1768b6c0000000000000000000000000000000000000000000000000000000000000002"
  //         });
  //         web31.eth.sendTransaction({
  //           from: account2.address,
  //           to: contract_address1,
  //           gas: 215650,
  //           data: "0xf1768b6c0000000000000000000000000000000000000000000000000000000000000002"
  //         });
  //       }
  //     })
  //   }
  //   else {
  //     log.consoleLog(workerId, "n....."+blc + " "+blc1)
  //     web3.eth.sendTransaction({
  //       from: account.address,
  //       to: contract_address,
  //       gas: 215650,
  //       data: "0xf1768b6c0000000000000000000000000000000000000000000000000000000000000002"
  //     });
  //     web31.eth.sendTransaction({
  //       from: account2.address,
  //       to: contract_address1,
  //       gas: 215650,
  //       data: "0xf1768b6c0000000000000000000000000000000000000000000000000000000000000002"
  //     });
  //   }
  // });



//initial
  // var ethContractDeployment = ethContract
  //   .deploy({ data: '0x' + bytecode, arguments: [60000000000000] })
  //   .send({ from: account.address, gas: 1000000 })
  //   .then(contract => {
  //     log.consoleLog(
  //       workerId,
  //       'Deployed contract address: ' + contract.options.address
  //     )
  //     web3.eth.getTransactionCount(account.address).then(nonce => {
  //       log.consoleLog(workerId, nonce)
  //       contract.methods
  //       .givetoothers('0x94abee207bc81159831e3c6b3a6722d019e8e13e', 0)
  //       .send({ from: account.address, gas: 1000000, 
  //       nonce: nonce, gasprice: (0 + 1)*200000 })
  //       .on('receipt', recp =>
  //       log.consoleLog(workerId, 'Transaction ' + 0 + ' injected.')
  //       )
  //       // transactionInjector(
  //       //   web3,
  //       //   contract,
  //       //   account,
  //       //   log,
  //       //   workerId,
  //       //   testDuration,
  //       //   contractLoad,
  //       //   nonce,
  //       //   0
  //       // )
  //     })
  //   })
  // .catch(() =>
  //     log.consoleLog(workerId, "Contract deployment failed."));
}

function main (argv) {
  args = require('minimist')(argv.slice(2))
  if (args['c'] !== undefined && !Array.isArray(args['c'])) {
    args['c'] = [args['c']]
  }

  if (
    args['help'] ||
    args['h'] ||
    args['t'] === undefined ||
    isNaN(args['t']) ||
    args['c'] === undefined ||
    !args['c'].every(i => !isNaN(i))
  ) {
    console.log(HELP_MSG)
    return
  }

  // var log = require('./TG_log')
  // const rlp = require('rlp');
  // const keccak = require('keccak');

  var nonce = args['t']

  // account1 = accounts[parseInt(0)].address

  // var sender = account1; //Requires a hex string as input!
  // log.consoleLog(0, 'nonce is ' + nonce)

  // var input_arr = [ sender, nonce ];
  // var rlp_encoded = rlp.encode(input_arr);

  // var contract_address_long = keccak('keccak256').update(rlp_encoded).digest('hex');

  // var contract_address = contract_address_long.substring(24);

  // log.consoleLog(0, 'contract address is ' + contract_address)
  worker(nonce)
}

main(process.argv)
// worker(process.argv)
