# adv-top-proj

Here is the `README.md` file for your Docker image setup instructions:

---

## Docker Image Setup Instructions

Follow these steps to run the backend and frontend Docker containers.

### 1. Pull Backend Image
Download the backend image from the Docker registry:
```bash
docker pull wishmarukapitak/backend:latest
```

### 2. Pull Frontend Image
Download the frontend image from the Docker registry:
```bash
docker pull wishmarukapitak/frontend:latest
```

### 3. Create a Directory for Backend
Create a directory named `backend` and place your `.env` file inside this directory. The `.env` file contains the environment variables required for the backend to function.

### 4. Run Backend with `.env`
Run the backend container, binding the `.env` file to the container:
```bash
docker run -p 8010:8010 -v $(pwd)/backend/.env:/app/.env wishmarukapitak/backend:latest
```

### 5. Run Frontend
Run the frontend container, exposing it on port 3000:
```bash
docker run -p 3000:80 wishmarukapitak/frontend:latest
```

---

### Notes:
- Ensure you have Docker installed on your machine.
- Replace the `.env` file content with the appropriate configuration for your backend service.
- Access the backend via `http://localhost:8010` and the frontend via `http://localhost:3000`.

Let me know if you need additional modifications!