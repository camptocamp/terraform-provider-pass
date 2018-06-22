import { moduleFor, test } from 'ember-qunit';

moduleFor('adapter:console', 'Unit | Adapter | console', {
  needs: ['service:auth', 'service:flash-messages', 'service:version'],
});

test('it builds the correct URL', function(assert) {
  let adapter = this.subject();
  let sysPath = 'sys/health';
  let awsPath = 'aws/roles/my-other-role';
  assert.equal(adapter.buildURL(sysPath), '/v1/sys/health');
  assert.equal(adapter.buildURL(awsPath), '/v1/aws/roles/my-other-role');
});
