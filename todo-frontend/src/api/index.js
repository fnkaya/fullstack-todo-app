import axios from 'axios';

export class API {
  constructor(url) {
    if (url === undefined || url === '') {
      // eslint-disable-next-line no-param-reassign
      url = process.env.VUE_APP_BASE_API_URL;
    }

    if (url.endsWith('/')) {
      // eslint-disable-next-line no-param-reassign
      url = url.substr(0, url.length - 1);
    }
    this.url = url;
  }

  withPath(path) {
    if (!path.startsWith('/')) {
      // eslint-disable-next-line no-param-reassign
      path = `/${path}`;
    }

    return `${this.url}${path}`;
  }

  getTodos() {
    return axios.get(this.withPath('/api/todos')).then((response) => response.data);
  }

  addTodo(data) {
    return axios.post(this.withPath('/api/todos'), data).then((response) => response.status);
  }
}

export default new API(process.env.VUE_APP_BASE_API_URL);
