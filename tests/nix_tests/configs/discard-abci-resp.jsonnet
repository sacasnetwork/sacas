local config = import 'default.jsonnet';

config {
  'sac_1312-1'+: {
    config+: {
      storage: {
        discard_abci_responses: true,
      },
    },
  },
}
