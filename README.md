# CS4035D: Computer Security

## Lamport Hash

The project is build using the GO programming language. Therefore GO should be installed to compile the project.

### Lamport Hash Backend

> Note: Once the backend is terminated all data will be lost.

To compile the backend, go to the backend folder.

```
cd backend
```

Run the command

```
go build .
```

A binary file with the name l-hash-backend will be created.

Run the below command to start the backend

```
./l-hash-backend
```

### Lamport Hash Frontend

To compile the frontend, go to the frontend folder.

```
cd frontend
```

Run the command

```
go build .
```

A binary file with the name l-hash-frontend will be created.

Run the below command to show the help page

```
./l-hash-frontend
```

> Note : The following commands requires one instance of the backend running. Otherwise requests will fail.

To register a user

```
./l-hash-frontend register
```

To Login as a user

```
./l-hash-frontend login
```

To reset the password of a user

```
./l-hash-frontend reset
```
