local config = import 'default.jsonnet';

config {
  'sacas_1317-1'+: {
    validators: super.validators[0:1] + [{
      name: 'fullnode',
    }],
    'app-config'+: {
      'api'+: {
        'enable': true,
      },
      'grpc'+: {
        'enable': true,
      },
    },
  },
}
