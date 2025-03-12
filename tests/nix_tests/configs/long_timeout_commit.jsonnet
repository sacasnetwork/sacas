local default = import 'default.jsonnet';

default {
  'sacas_1317-1'+: {
    config+: {
      consensus+: {
        timeout_commit: '5s',
      },
    },
  },
}
