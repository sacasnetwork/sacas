local config = import 'default.jsonnet';

config {
  'sacas_1317-1'+: {
    config+: {
      storage: {
        discard_abci_responses: true,
      },
    },
  },
}
