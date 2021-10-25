## HyperService 03-01

主要就三个Repo: NSB, ves, uip

+ ves已经完全重写了

  + 原来的项目存在一些设计不合理的地方
  + 并不存在难以修复的bug，大量代码仍然基于原项目
  + 主要增加了相当多的测试
    + 任何稳定的项目都需要足够多的测试覆盖代码的每一个细节，但是没有人力没有时间，现在不可能做得太完美了

+ uip

  + 修改了op-intent的结构，添加了if和loop的parse。

  ```plain
  ├─op-intent
  │  ├─document
  │  ├─errorn
  │  ├─lexer // js的json转为golang的struct，再统一变成方便的形式
  │  ├─parser // 生成平坦的指令流
  │  └─token
  ```

  js给我的json

  ```json
  {
    "op-intents": [
      {
        "name": "op1",
        "type": "Payment",
        "src": {
          "domain": 1,
          "user_name": "a1"
        },
        "dst": {
          "domain": 2,
          "user_name": "a2"
        },
          // 修改成了常数，amount应该是容忍的最大值，不是确定值
        "amount": {
          "type": "uint256",
          "value": {
            "contract": "c2",
            "field": "totalVotes",
            "pos": "01"
          }
        },
        "unit": "ether",
        "ratio": "1 ether as XYZ"
      },
      {
        "name": "op2",
        "type": "ContractInvocation",
        "invoker": "a2",
        "func": "vote",
        "contract": {
          "domain": 2,
          "address": "0x3723261b2a5a62b778b5c74318534d7fdf8db38c"
        },
        "parameters": []
      },
      {
        "name": "if-op",
        "type": "IfStatement",
        "if": [
          {
            "name": "op3",
            "type": "ContractInvocation",
            "invoker": "a2",
            "func": "vote",
            "contract": {
              "address": "0x3723261b2a5a62b778b5c74318d34d7fdbadb38e"
            },
            "parameters": []
          },
          {
            "name": "op4",
            "type": "Payment",
            "src": {
              "domain": 1,
              "user_name": "a1"
            },
            "dst": {
              "domain": 2,
              "user_name": "a2"
            },
            "amount": "aa",
            "unit": "ether",
            "ratio": "1 ether as XYZ"
          }
        ],
        "else": [
          {
            "name": "op5",
            "type": "ContractInvocation",
            "invoker": "a2",
            "func": "vote",
            "contract": {
              "domain": 2,
              "address": "0x3723261b2a5a62b778b5c74318534d7fdf8db38c"
            },
            "parameters": []
          }
        ],
        "condition": {
          "left": {
            "type": "uint256",
            "value": {
              "contract": "c2",
              "field": "num_count",
              "pos": "00"
            }
          },
          "right": {
            "type": "uint256",
            "value": {
              "contract": "c2",
              "field": "totalVotes",
              "pos": "01"
            }
          },
          "sign": "Greater"
        }
      },
      {
        "name": "loop",
        "type": "loopFunction",
        "loop": [
          {
            "name": "op6",
            "type": "ContractInvocation",
            "invoker": "a2",
            "func": "vote",
            "contract": {
              "domain": 2,
              "address": "0x3723261b2a5a62b778b5c74318534d7fdf8db38c"
            },
            "parameters": []
          }
        ],
        "loopTime": "5"
      }
    ],
    "dependencies": [
      {
        "left": "Op1",
        "right": "Op2",
        "dep": "before"
      }
    ],
    "contracts": [
      {
        "contractName": "c1",
        "domain": 1,
        "address": "0xafc7d2959e72081770304f6474151293be1fbba7"
      },
      {
        "contractName": "c2",
        "domain": 2,
        "address": "0x3723261b2a5a62b778b5c74318534d7fdf8db38c"
      },
      {
        "contractName": "c3",
        "domain": 3,
        "address": "0x3723261b2a5a62b778b5c74318d34d7fdbadb38e"
      }
    ],
    "accounts": [
      {
        "userName": "a1",
        "domain": 1,
        "address": "0x7019fa779024c0a0eac1d8475733eefe10a49f3b"
      },
      {
        "userName": "a2",
        "domain": 2,
        "address": "0x47a1cdb6594d6efed3a6b917f2fbaa2bbcf61a2e"
      },
      {
        "userName": "a3",
        "domain": 3,
        "address": "0x47a1cdb6559d6efed3a6b917f2fbaa2bbcf61a2e"
      }
    ]
  }
  ```

  编译成指令流:

  ```go
  op1.cna
  op1.cnb
  op2
  if-op.goto.if
  op5
  if-op.goto.else
  op3
  op4.cna
  op4.cnb
  loop.loopBegin
  op6
  loop.addLoopVar
  loop.loopEnd
  loop.resetLoopVar
  ```

  编译方法:

  ```go
  // if cond [if-op]
  //    op1
  //    op2
  //    ...
  //    opk
  // else
  //    opk+1
  //    opk+2
  //    ...
  //    opk+n
  
  // ->
  
  // goto index(op1) if cond
  // opk+1
  // opk+2
  // ...
  // opk+n
  // goto index(opk)+1 if true
  // op1
  // op2
  // ...
  // opk
  ```

  ```go
  // loop times
  //    op1
  //    op2
  //    ...
  //    opk
  
  // ->
  
  // loopBegin: goto reset if loopVar >= times
  // op1
  // op2
  // ...
  // opk
  // loopVar := loopVar + 1
  // loopEnd: goto loopBegin if true
  // reset: loopVar := 0
  ```

  

  + 修改了isc的内容，支持在线的eval和pc跳转

    + eval支持常量和isc内部的变量，暂时不支持state

    ```go
    
    func NewISC(msg Context, storage *storage.VM) *ISC {}
    
    type ISC struct {
            Storage Storage // localstorage
            Msg     Context // msg.sender, msg.value, 等等
    }
    
    func (isc *ISC) FreezeInfo(tid uint64) Response {}
    
    func (isc *ISC) maybeSwitchToStateSettling(pc uint64) Response {}
    
    func (isc *ISC) UserAck(fr, signature []byte) Response {
    	assertTrue(isc.IsInitialized(), CodeIsNotInitialized)
    	acknowledged := isc.Storage.UserAcknowledged()
    	if acknowledged.Get(fr) == nil {
    		acknowledged.Set(fr, signature)
    		uac := isc.Storage.getUserAckCount() + 1
    		isc.Storage.setUserAckCount(uac)
            
            // 当所有人都认可ves的初始化时
    		if uac == isc.Storage.Owners().Length() {
                // 初始化 pc
    			pc := isc.initPC(0)
    			isc.Storage.setPC(pc)
    			isc.Storage.setISCState(StateOpening)
    			if r := isc.maybeSwitchToStateSettling(pc); r != nil {
    				return r
    			}
    		}
    	}
    
    	return OK
    }
    
    func (isc *ISC) InsuranceClaim(tid, aid uint64) Response {
    	assertTrue(isc.IsInitialized(), CodeIsNotOpening)
    	var pc = isc.Storage.getPC()
    	assertTrue(pc == tid, CodeTransactionNotActive)
        
        // miuPC，即所有transaction的运行记录暂时是个问题，暂时先不统计。先能运行再说
    	var AidMap = isc.Storage.AidMap()
    	var miuPC = AidMap.Get(tid) + 1
    	assertTrue(miuPC == aid, CodeActionNotActive)
    	AidMap.Set(pc, miuPC)
    	if miuPC == TxState.Closed {
            // 重新计算 pc 的位置
    		pc = isc.nextPC(pc)
    		isc.Storage.setPC(pc)
            if r := isc.maybeSwitchToStateSettling(pc); r != nil {
                return r
            }
    	}
    }
    
    func (isc *ISC) SettleContract() Response {
    	assertTrue(isc.IsSettling(), CodeIsNotSettling)
        // 计算isc的成本，归还给大家多少钱
    	isc.Storage.setISCState(StateClosed)
    	return OK
    }
    
    func (isc *ISC) initPC(pc uint64) (uint64, error) {
    	instruction := opintent.HandleInstruction(isc.Storage.Instructions().Get(pc))
    	switch instruction.GetType() {
    	case instruction_type.Payment, instruction_type.ContractInvoke:
    		return pc, nil
            // pc要移动到transaction类型的instruction
    	default:
    		return isc._nextPC(pc, instruction)
    	}
    }
    
    func (isc *ISC) nextPC(pc uint64) (uint64, error) {
    	instruction := opintent.HandleInstruction(isc.Storage.Instructions().Get(pc))
    	return isc._nextPC(pc, instruction)
    }
    
    func (isc *ISC) _nextPC(pc uint64, instruction opintent.LazyInstruction) (uint64, error) {
    	switch instruction.GetType() {
            //无条件跳转
    	case instruction_type.Goto:
    		i := instruction.Deserialize()
    		return isc.nextPC(uint64(i.(*parser.Goto).Index))
            //有条件跳转
    	case instruction_type.ConditionGoto:
    		...
            //计算条件
    		if isc.evalBytes(
                i.(*parser.ConditionGoto).Condition).GetConstant().(bool) {
    			return isc.nextPC(uint64(i.(*parser.ConditionGoto).Index))
    		}
    		return isc.nextPC(pc+1)
            //有条件修改isc的变量
    	case instruction_type.ConditionSetState:
    		...
            //无条件修改isc的变量
    	case instruction_type.SetState:
    		...
    	default:
    		return isc.nextPC(pc+1)
    	}
    }
    
    func (isc *ISC) eval(v token.Param) (token.ConstantI, error) {
    	switch v.GetType(){
            //常量
    	case token.Constant:
    		return v.(token.ConstantI), nil
            //isc本地的变量
    	case token.LocalStateVariable:
    		v := v.(token.LocalStateVariableI)
    		return isc.load(v.GetField(), v.GetParamType())
            // 二元运算
    	case token.BinaryExpression:
    		v := v.(token.BinaryExpressionI)
    		l := isc.eval(v.GetLeft())
    		r := isc.eval(v.GetRight())
    		switch v.GetSign() {
    		case sign_type.EQ:
    			return eq(l, r)
    		case sign_type.LE:
    			return le(l, r)
    		case sign_type.LT:
    			...
    		default:
    			return nil, fmt.Errorf("unknown sign_type: %v", v.GetSign())
    		}
            //一元运算
    	case token.UnaryExpression:
    		v := v.(token.BinaryExpressionI)
    		l := isc.eval(v.GetLeft())
    		switch v.GetSign() {
    		case sign_type.LNot:
    			return lnot(l)
    		}
            //计算任意blockchain上的状态变量
    	case token.StateVariable:
    		//v := v.(token.StateVariableI)
    		return nil, errors.New("todo")
    	}
    	return nil, errors.New("param type not found")
    }
    
    func (isc *ISC) save(variable token.ConstantI) ([]byte, error) {
    	return storage.Encode(variable)
    }
    
    func (isc *ISC) load(field []byte, t value_type.Type) (token.ConstantI, error) {
    	return storage.Decode(isc.Storage.storage.GetBytes(string(field)), t)
    }
    
    ```

+ 剩下的工作

  + 将isc的测试写了
  + 把isc接到nsb上
  + ves预计几乎不需要修改，先做一轮测试再看情况分析

+ 代码统计，代码覆盖率暂时不放上来了等测试完了再放上来

  + uip

  ```plain
  C:\work\go\src\github.com\HyperService-Consortium\go-uip (master -> origin)
  λ cloc .
       245 text files.
       243 unique files.
        70 files ignored.
  
  github.com/AlDanial/cloc v 1.84  T=0.50 s (454.0 files/s, 66654.0 lines/s)
  -------------------------------------------------------------------------------
  Language                     files          blank        comment           code
  -------------------------------------------------------------------------------
  Go                             140           1668            971           9892
  -------------------------------------------------------------------------------
  SUM:                           227           3723           3883          25721
  -------------------------------------------------------------------------------
  
  
  ```

  + ves

  ```plain
  C:\work\go\src\github.com\HyperService-Consortium\go-ves (master -> origin)      
  λ cloc .                                                                         
       603 text files.                                                             
       578 unique files.                                                           
       189 files ignored.                                                          
                                                                                   
  github.com/AlDanial/cloc v 1.84  T=0.50 s (1114.0 files/s, 84228.0 lines/s)      
  -------------------------------------------------------------------------------- 
  Language                      files          blank        comment           code 
  -------------------------------------------------------------------------------- 
  Go                              504           5001           2380          31004 
  Python                           16            584            110           2126 
  Protocol Buffers                  3             76             35            243 
  -------------------------------------------------------------------------------- 
  SUM:                            557           5716           2537          33861 
  -------------------------------------------------------------------------------- 
  ```

  + NSB, 不是这个寒假写的, 几乎没有变化