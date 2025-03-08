{
  dotenv: '../../../scripts/.env',
  'sac_1312-1': {
    'account-prefix': 'sacas',
    'coin-type': 60,
    cmd: 'evmosd',
    'start-flags': '--trace',
    'app-config': {
      'app-db-backend': 'goleveldb',
      'minimum-gas-prices': '0sac',
      'index-events': ['ethereum_tx.ethereumTxHash'],
      'json-rpc': {
        address: '127.0.0.1:{EVMRPC_PORT}',
        'ws-address': '127.0.0.1:{EVMRPC_PORT_WS}',
        api: 'eth,net,web3,debug',
        'feehistory-cap': 100,
        'block-range-cap': 10000,
        'logs-cap': 10000,
        'fix-revert-gas-refund-height': 1,
        enable: true,
      },
      api: {
        enable: true
      }
    },
    validators: [{
      coins: '10001000000000000000000sac',
      staked: '1000000000000000000sac',
      mnemonic: '${VALIDATOR1_MNEMONIC}',
    }, {
      coins: '10001000000000000000000sac',
      staked: '1000000000000000000sac',
      mnemonic: '${VALIDATOR2_MNEMONIC}',
    }],
    accounts: [{
      name: 'community',
      coins: '10000000000000000000000sac',
      mnemonic: '${COMMUNITY_MNEMONIC}',
    }, {
      name: 'signer1',
      coins: '20000000000000000000000sac',
      mnemonic: '${SIGNER1_MNEMONIC}',
    }, {
      name: 'signer2',
      coins: '30000000000000000000000sac',
      mnemonic: '${SIGNER2_MNEMONIC}',
    }],
    genesis: {
      consensus_params: {
        block: {
          max_bytes: '1048576',
          max_gas: '81500000',
        },
      },
      app_state: {
        evm: {
          params: {
            evm_denom: 'sac',
          },
        },
        crisis: {
          constant_fee: {
            denom: 'sac',
          },
        },
        staking: {
          params: {
            bond_denom: 'sac',
          },
        },
        inflation: {
          params: {
            mint_denom: 'sac',
          },
        },
        gov: {
          deposit_params: {
            max_deposit_period: '10s',
            min_deposit: [
              {
                denom: 'sac',
                amount: '1',
              },
            ],
          },
          params: {
            min_deposit: [
              {
                denom: 'sac',
                amount: '1',
              },
            ],
            max_deposit_period: '10s',
            voting_period: '10s',         
            expedited_voting_period: '5s',   
          },
        },
        transfer: {
          params: {
            receive_enabled: true,
            send_enabled: true,
          },
        },
        feemarket: {
          params: {
            no_base_fee: false,
            base_fee: '100000000000',
          },
        },
      },
    },
  },
}
