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

//gaslimit initially 0x10000000

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



function worker () {
  
  workerId = 0

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

  log.consoleLog(workerId, 'Ethereum account ' + account.address + ' used.')
  log.consoleLog(workerId, 'Ethereum account private key ' + account.privateKey)

  


  log.consoleLog(workerId, 'compling...')


  var source = fs.readFileSync('./twofunctions.sol', 'utf8');
  var input = {
    language: 'Solidity',
    sources: {
      'twofunctions.sol': {
        content: source
      }
    },
    settings: {
      outputSelection: {
        '*': {
          '*': ['*']
        }
      }
    }
  };
  var solcout = JSON.parse(solc.compile(JSON.stringify(input)));

  
  // log.consoleLog(workerId, JSON.stringify(solcout))
  log.consoleLog(workerId, 'Completed compiling.');


  log.consoleLog(workerId, 'test............')


  for (var contractName in solcout.contracts['twofunctions.sol']) {
    log.consoleLog(workerId,
      contractName +
        ': ' +
        solcout.contracts['twofunctions.sol'][contractName].evm.bytecode.object
    );
    // log.consoleLog(workerId,
    //   contractName + ':' + JSON.stringify((solcout.contracts['twofunctions.sol']['HelloWorld'].abi))
    // );
  }
  

  var abi = (solcout.contracts['twofunctions.sol']['HelloWorld'].abi)
  var bytecode = solcout.contracts['twofunctions.sol'][contractName].evm.bytecode.object

  // log.consoleLog(workerId, bytecode)

  var ethContract = new web3.eth.Contract(abi)


  // log.consoleLog(workerId, web3.utils.sha3("update(uint256)"));
  //update(uint): 0x82ab890a4924aa641094939d7f06fdb5d410dc84a4205ffbb6c20dfc50e7f984
  //givetoothers(): 0xfdc4b27ecbcc5f82e2dee841ed80ee1224599ab840f05fe7180be0ab48508281
  //givetoothers(): 0xfdc4b27e
  //givetoinitial(): 0x202ec152a994b8eb829ddb591f8ccaa7a603ccaab3f520846be7f0bba880c123
  var data1 = "0xf1768b6c"
  var data2 = "0000000000000000000000000000000000000000000000000000000000000001"
  var datatotal = "0xf1768b6c0000000000000000000000000000000000000000000000000000000000000001"

  // log.consoleLog(workerId, web3.utils.sha3("givetoinitial()"))
  //var datagivetoothers = "0xda7d2982000000000000000000000000b782edeacc48fa6161dc690043bd0db17e959bf4"
  var datagivetoothers = "0xda7d2982000000000000000000000000ab8483f64d9c6d1ecf9b849ae677dd3315835cb2"
  var datagivetoinitial = "0x202ec152"

  // log.consoleLog(workerId, datatotal)


  const rlp = require('rlp');
  const keccak = require('keccak');

  var nonce = 0x00; //The nonce must be a hex literal!
  var sender = account.address; //Requires a hex string as input!

  var input_arr = [ sender, nonce ];
  var rlp_encoded = rlp.encode(input_arr);

  var contract_address_long = keccak('keccak256').update(rlp_encoded).digest('hex');

  var contract_address = contract_address_long.substring(24); //Trim the first 24 characters.
  // var contract_address = "0xD6937b20Fb09568094aa69BDd666d2266548130d"
  log.consoleLog(workerId, "contract_address: " + contract_address);

// //0xD6937b20Fb09568094aa69BDd666d2266548130d
// //0xD6937b20Fb09568094aa69BDd666d2266548130d
// //0xd6937b20fb09568094aa69bdd666d2266548130d

  web3.eth.sendTransaction({
    from: account.address,
    gas: 398928,
    value: 3000000000000000000,
    data: '0x'+ bytecode
  }, function(error, hash) {
//to be added
  }).on('receipt', recp=>{
    
  });

//   web3.eth.sendTransaction({
//     from: account.address, 
//     to: contract_address,
//     gas: 21565,
//     data: datagivetoothers
//   }, function(error, hash) {
// //tod 
//   });

  // var account1 = web3.eth.accounts.privateKeyToAccount(accounts[1].privateKey);
  // web3.eth.accounts.wallet.add(account1)

  // web3.eth.sendTransaction({
  //   from: account.address, 
  //   to: contract_address,
  //   gas: 215650,
  //   data: datagivetoinitial
  // }, function(error, hash){
  //   //to ...
  // });

  // web3.eth.estimateGas({
  //   from: account.address,
  //   // gas: 1000000,
  //   value: 1*10**18,
  //   data: '0x'+ bytecode // deploying a contracrt
  // }).then(gasa=>log.consoleLog(workerId, gasa));

  



  // web3.eth.sendTransaction({
  //   from: account.address,
  //   gas: 335000,
  //   // value: 60000000000000,
  //   data: '0x'+ bytecode // deploying a contracrt
  // }).on('receipt', recp=>{
  //   web3.eth.estimateGas({
  //     from: account.address,
  //     to: contract_address,
  //     // value: 60000000000000,
  //     data: datatotal // 
  //   }).then(gasa=>log.consoleLog(workerId, gasa));

  // })

  // // var data = "608060405260405161044a38038061044a83398181016040528101906100259190610090565b8060008190555033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600181905550506100e3565b60008151905061008a816100cc565b92915050565b6000602082840312156100a6576100a56100c7565b5b60006100b48482850161007b565b91505092915050565b6000819050919050565b600080fd5b6100d5816100bd565b81146100e057600080fd5b50565b610358806100f26000396000f3fe60806040526004361061003f5760003560e01c8063202ec1521461004457806382ab890a1461004e578063b69ef8a814610077578063da7d2982146100a2575b600080fd5b61004c6100be565b005b34801561005a57600080fd5b506100756004803603810190610070919061025c565b610140565b005b34801561008357600080fd5b5061008c61019f565b6040516100999190610298565b60405180910390f35b6100bc60048036038101906100b7919061022f565b6101a5565b005b6002600154141561013e57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc620f42409081150290604051600060405180830381858888f19350505050158015610134573d6000803e3d6000fd5b5060066001819055505b565b6005600154148061015357506006600154145b8061015f575060018054145b8061016c57506002600154145b1561017f5760015460018190555061019c565b600181148061018e5750600281145b1561019b57806001819055505b5b50565b60005481565b600180541415610202578073ffffffffffffffffffffffffffffffffffffffff166108fc620f42409081150290604051600060405180830381858888f193505050501580156101f8573d6000803e3d6000fd5b5060056001819055505b50565b600081359050610214816102f4565b92915050565b6000813590506102298161030b565b92915050565b600060208284031215610245576102446102ef565b5b600061025384828501610205565b91505092915050565b600060208284031215610272576102716102ef565b5b60006102808482850161021a565b91505092915050565b610292816102e5565b82525050565b60006020820190506102ad6000830184610289565b92915050565b60006102be826102c5565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600080fd5b6102fd816102b3565b811461030857600080fd5b50565b610314816102e5565b811461031f57600080fd5b5056fea2646970667358221220d2708628f7a86bf3c8d04067be7b61370b67bfe069de1144e156cc817278b2bc64736f6c63430008070033"
  // web3.eth.sendTransaction({
  //   from: account.address,
  //   gas: 331372,
  //   // value: 60000000000000,
  //   data: '0x'+ bytecode // deploying a contracrt
  // }).on('receipt', recp=> {
  //   // log.consoleLog(workerId, recp)
  //   web3.eth.sendTransaction({
  //     from: account1.address, 
  //     to: contract_address,
  //     gas: 21565,
  //     data: datatotal
  //   }).on('receipt', recp=>{
  //     // web3.eth.sendTransaction({
  //     //   from: account1.address, 
  //     //   to: contract_address,
  //     //   gas: 21565,
  //     //   data: datagivetoothers
  //     // }).on('receipt', recp=>{
        
  //     // })
  //   });
  // });

  



  // web3.eth.estimateGas({
  //   from: account.address,
  //   to: contract_address,
  //   // value: 60000000000000,
  //   data: datatotal // 
  // }).then(gasa=>log.consoleLog(workerId, gasa));

  




//initial
  // var ethContractDeployment = ethContract
  //   .deploy({ data: '0x' + bytecode })
  //   .send({ from: account.address, gas: 335000 }, function(error, hash){})
  //   .then(contract => {
  //     log.consoleLog(
  //       workerId,
  //       'Deployed contract address: ' + contract.options.address
  //     )
  //     // web3.eth.getTransactionCount(account.address).then(nonce => {
  //     //   log.consoleLog(workerId, nonce)
  //     //   contract.methods
  //     //   .givetoothers('0x94abee207bc81159831e3c6b3a6722d019e8e13e', 0)
  //     //   .send({ from: account.address, gas: 1000000, 
  //     //   nonce: nonce, gasprice: (0 + 1)*200000 })
  //     //   .on('receipt', recp =>
  //     //   log.consoleLog(workerId, 'Transaction ' + 0 + ' injected.')
  //     //   )
  //     // })
  //   })
  // .catch(() =>
  //     log.consoleLog(workerId, "Contract deployment failed."));



}


//todo: hyperservice: web3, change to launching smart contracts..., write a script for gettointial, and send transactions to hyperservice; write a contract on hyperservice
function main (argv) {
  worker()
}

main(process.argv)
// worker(process.argv)
