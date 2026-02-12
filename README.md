# 📘 Documentação — CRUD de Tasks em Go (Gin)

## 🧭 Visão Geral

Este projeto implementa um CRUD completo utilizando a linguagem **Go** e o framework **Gin** para criação de uma API REST.

O objetivo é praticar:

* Estruturas (`struct`)
* Rotas HTTP
* Manipulação de JSON
* Slice como armazenamento
* Organização de código backend

Este projeto não utiliza banco de dados — os dados são mantidos em memória (ideal para aprendizado).

---

## 🛠 Tecnologias Utilizadas

* Go
* Gin Gonic (Framework HTTP)

---

## 📦 Instalação e Configuração

### 1️⃣ Criar o projeto

```bash
mkdir crud-go
cd crud-go
```

---

### 2️⃣ Inicializar módulo Go

```bash
go mod init crud-go
```

Isso cria o sistema de dependências do projeto.

---

### 3️⃣ Instalar Gin

```bash
go get github.com/gin-gonic/gin
```

---

### 4️⃣ Criar arquivo principal

```bash
touch main.go
```

---

## 🧱 Estrutura do Código

### ✔️ Struct Task

```go
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
```

Representa uma tarefa.

Campos:

* ID → Identificador único
* Title → Título da tarefa
* Description → Descrição
* Done → Status (concluída ou não)

`json:"campo"` permite conversão automática entre JSON e Go.

---

### ✔️ Armazenamento em memória

```go
var tasks = []Task{}
var nextID = 1
```

* `tasks` → slice que simula um banco
* `nextID` → gera IDs automáticos

---

## 🚀 Inicialização do Servidor

```go
r := gin.Default()
```

Cria servidor HTTP com:

* Logger
* Recovery de erros

---

### Rotas definidas

```go
r.POST("/tasks", createTask)
r.GET("/tasks", getTasks)
r.GET("/tasks/:id", getTaskByID)
r.PUT("/tasks/:id", updateTask)
r.DELETE("/tasks/:id", deleteTask)
```

Cada rota aponta para uma função (handler).

---

## 🔵 CREATE — Criar tarefa

### Função

```go
func createTask(c *gin.Context)
```

### Passo a passo

1️⃣ Recebe JSON da requisição

```go
c.ShouldBindJSON(&newTask)
```

2️⃣ Converte JSON → struct

3️⃣ Gera ID

```go
newTask.ID = nextID
nextID++
```

4️⃣ Adiciona ao slice

```go
tasks = append(tasks, newTask)
```

5️⃣ Retorna resposta

```go
c.JSON(http.StatusCreated, newTask)
```

---

## 🟢 READ — Listar tarefas

```go
func getTasks(c *gin.Context)
```

Simplesmente retorna o slice:

```go
c.JSON(http.StatusOK, tasks)
```

---

## 🟡 READ — Buscar por ID

```go
func getTaskByID(c *gin.Context)
```

### Etapas

1️⃣ Pega ID da URL

```go
id, _ := strconv.Atoi(c.Param("id"))
```

2️⃣ Percorre tarefas

```go
for _, t := range tasks
```

3️⃣ Se encontrar → retorna

4️⃣ Se não → erro 404

---

## 🟠 UPDATE — Atualizar tarefa

```go
func updateTask(c *gin.Context)
```

### Processo

1️⃣ Lê ID
2️⃣ Recebe novo JSON
3️⃣ Procura tarefa
4️⃣ Atualiza campos
5️⃣ Retorna tarefa atualizada

```go
tasks[i].Title = update.Title
```

---

## 🔴 DELETE — Remover tarefa

```go
func deleteTask(c *gin.Context)
```

### Etapas

1️⃣ Busca índice da tarefa
2️⃣ Remove do slice

```go
tasks = append(tasks[:i], tasks[i+1:]...)
```

Isso cria um novo slice sem o item deletado.

3️⃣ Retorna confirmação

---

## ▶️ Executando o Projeto

```bash
go run main.go
```

Servidor inicia em:

```
localhost:8080
```

---

## 🧪 Testando Endpoints

### Criar

POST `/tasks`

```json
{
"title": "Estudar Go",
"description": "Praticar CRUD",
"done": false
}
```

---

### Listar

GET `/tasks`

---

### Buscar

GET `/tasks/1`

---

### Atualizar

PUT `/tasks/1`

---

### Deletar

DELETE `/tasks/1`

---

## 📚 Conceitos Aprendidos

* REST API
* JSON Binding
* Struct Tags
* Slices
* Parâmetros de rota
* Organização backend básica

---

## 🚀 Possíveis Evoluções

* Persistência com PostgreSQL
* Arquitetura em camadas
* Autenticação JWT
* Docker
* Testes automatizados
* Logs estruturados

---

## 🎯 Conclusão

Este projeto demonstra os fundamentos de construção de APIs em Go e serve como base para aplicações backend reais. Ele prepara o desenvolvedor para evoluir para arquiteturas mais complexas e integração com banco de dados.
