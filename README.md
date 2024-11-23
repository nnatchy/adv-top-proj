# adv-top-proj

```bash
By Group mairoira,
6431307421 Jirawat Lengnoi
6431318321 Natthasith Wiriyayothin
6431342321 Wish Marukapitak
```

---
Here’s the revised `README.md` with the additional note about checking the `README.md` files in the respective directories for more information if something isn’t working:

---

# Docker Image Setup Instructions

Follow these steps to run the backend and frontend Docker containers and ensure they can communicate with each other.

---

## 1. Pull Backend Image
Download the backend image from the Docker registry:
```bash
docker pull wishmarukapitak/backend:latest
```

## 2. Pull Frontend Image
Download the frontend image from the Docker registry:
```bash
docker pull wishmarukapitak/frontend:latest
```

---

## 3. Create a Custom Docker Network
Create a custom Docker network to enable communication between the backend and frontend containers:
```bash
docker network create my_custom_network
```

---

## 4. Create a Directory for Backend
Create a directory named `backend` and place your `.env` file inside this directory. The `.env` file contains the environment variables required for the backend to function.

---

## 5. Run Backend with `.env`
Run the backend container and attach it to the custom network:
```bash
docker run --network my_custom_network -p 8010:8010 -v $(pwd)/backend/.env:/app/.env wishmarukapitak/backend:latest
```

---

## 6. Run Frontend
Run the frontend container and attach it to the same custom network:
```bash
docker run --network my_custom_network -p 3000:80 wishmarukapitak/frontend:latest
```

---

### Notes:
1. **Communication Between Backend and Frontend**:
   - Ensure the backend and frontend are configured to communicate within the same network. For example, in the `.env` file or frontend configuration, reference the backend service using the container name (e.g., `http://backend:8010`).
   
2. **Access the Services**:
   - Access the backend via `http://localhost:8010`.
   - Access the frontend via `http://localhost:3000`.

3. **For More Information**:
   - If something isn’t working properly, check the respective `README.md` file in the `frontend` or `backend` directory for more details and troubleshooting steps.

---

### Troubleshooting
- **Network Issues**:
  - Ensure both containers are in the same network using the following command:
    ```bash
    docker network inspect my_custom_network
    ```
- **Environment Variables**:
  - Verify that the `.env` file contains the correct settings for your backend service.
  
---

This ensures that users know to reference the additional `README.md` files in each directory for more detailed instructions or debugging help. Let me know if you need further updates!