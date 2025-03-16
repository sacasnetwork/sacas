local config = import 'default.jsonnet';

config {
  'sacas_1317-1'+: {
    cmd: 'sacasd-rocksdb',    
    'app-config'+: {
      'app-db-backend': 'rocksdb',      
      pruning: 'everything',
      'state-sync'+: {
        'snapshot-interval': 0,
      },
      'memiavl'+:{
        enable: true,
        'snapshot-keep-recent': 0,
        'snapshot-interval': 1,
      },
      'versiondb'+: {
        enable: false,
      },
    },
    config+: {
       'db_backend': 'rocksdb',
    },
    genesis+: {
      app_state+: {
        feemarket+: {
          params+: {
            min_gas_multiplier: '0',
          },
        },
      },
    },
  },
}
