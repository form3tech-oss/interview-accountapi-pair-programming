import Ember from 'ember';
const ACTIONS = {
  performance: {
    primary: ['disable', 'demote', 'recover', 'reindex'],
    secondary: ['disable', 'promote', 'update-primary', 'recover', 'reindex'],
    bootstrapping: ['disable', 'recover', 'reindex'],
  },
  dr: {
    primary: ['disable', 'recover', 'reindex', 'demote'],
    secondary: ['promote'],
    bootstrapping: ['disable', 'recover', 'reindex'],
  },
};

export function replicationActionForMode([replicationMode, clusterMode] /*, hash*/) {
  return Ember.get(ACTIONS, `${replicationMode}.${clusterMode}`);
}

export default Ember.Helper.helper(replicationActionForMode);
