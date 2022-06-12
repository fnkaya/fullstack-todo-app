import { pactWith } from 'jest-pact';
import { eachLike, like } from '@pact-foundation/pact/src/dsl/matchers';
// eslint-disable-next-line import/no-named-as-default
import { API } from '@/api';

const getRequest = {
  method: 'GET',
  path: '/api/todos',
};

const getResponse = {
  status: 200,
  headers: {
    'Content-Type': 'application/json; charset=UTF-8',
  },
  body: eachLike({
    id: like(1),
    text: like('todo text'),
    done: like(false),
  }),
};

pactWith({ consumer: 'todos-fe', provider: 'todos-be' }, (provider) => {
  describe('todos contract tests', () => {
    let api;
    beforeEach(() => {
      api = new API(provider.mockService.baseUrl);
    });

    it('fetch todos', async () => {
      await provider.addInteraction({
        state: 'fetch todos successfully',
        uponReceiving: 'not empty todo list',
        withRequest: getRequest,
        willRespondWith: getResponse,
      });

      const response = await api.getTodos();
      expect(response[0].id).toEqual(1);
    });
  });
});
