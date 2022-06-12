```
helm package todo-backend -d ~/chart-museum/docs
helm repo index docs --url https://fnkaya.github.io/chart-museum/

helm repo add todo-backend-repo https://fnkaya.github.io/chart-museum/
helm repo update
helm upgrade --install todo-backend todo-backend-repo/todo-backend --version 0.3.0 --set image=fnkaya/todo-http-server:f3036e19
```