import Ember from 'ember';
import DS from 'ember-data';
import fetch from 'fetch';

const POLLING_URL_PATTERNS = ['sys/seal-status', 'sys/health', 'sys/replication/status'];

export default DS.RESTAdapter.extend({
  auth: Ember.inject.service(),

  flashMessages: Ember.inject.service(),

  namespace: 'v1/sys',

  shouldReloadAll() {
    return true;
  },

  shouldReloadRecord() {
    return true;
  },

  shouldBackgroundReloadRecord() {
    return false;
  },

  _preRequest(url, options) {
    const token = this.get('auth.currentToken');
    if (token && !options.unauthenticated) {
      options.headers = Ember.assign(options.headers || {}, {
        'X-Vault-Token': token,
      });
      if (options.wrapTTL) {
        Ember.assign(options.headers, { 'X-Vault-Wrap-TTL': options.wrapTTL });
      }
    }
    const isPolling = POLLING_URL_PATTERNS.some(str => url.includes(str));
    if (!isPolling) {
      this.get('auth').setLastFetch(Date.now());
    }
    if (this.get('auth.shouldRenew')) {
      this.get('auth').renew();
    }
    options.timeout = 60000;
    return options;
  },

  ajax(url, type, options = {}) {
    let opts = this._preRequest(url, options);

    return this._super(url, type, opts).then((...args) => {
      const [resp] = args;
      if (resp && resp.warnings) {
        const flash = this.get('flashMessages');
        resp.warnings.forEach(message => {
          flash.info(message);
        });
      }
      return Ember.RSVP.resolve(...args);
    });
  },

  // for use on endpoints that don't return JSON responses
  rawRequest(url, type, options = {}) {
    let opts = this._preRequest(url, options);
    return fetch(url, {
      method: type | 'GET',
      headers: opts.headers | {},
    }).then(response => {
      if (response.status >= 200 && response.status < 300) {
        return Ember.RSVP.resolve(response);
      } else {
        return Ember.RSVP.reject();
      }
    });
  },

  handleResponse(status, headers, payload, requestData) {
    const returnVal = this._super(...arguments);
    // ember data errors don't have the status code, so we add it here
    if (returnVal instanceof DS.AdapterError) {
      Ember.set(returnVal, 'httpStatus', status);
      Ember.set(returnVal, 'path', requestData.url);
    }
    return returnVal;
  },
});
