local default = import 'default.jsonnet';

default {
  'sac_1312-1'+: {
    config+: {
      consensus+: {
        timeout_commit: '5s',
      },
    },
  },
}
