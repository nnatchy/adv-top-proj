# adv-top-proj

```bash
By Group mairoira,
6431307421 Jirawat Lengnoi
6431318321 Natthasith Wiriyayothin
6431342321 Wish Marukapitak
```

---

# Docker Image Setup Instructions

Follow these steps to run the backend and frontend Docker containers and ensure they can communicate with each other.

---

## 1. Pull Backend Image
Download the backend image from the Docker registry:
```bash
docker pull nnatchy14/frontend:latest
```

## 2. Pull Frontend Image
Download the frontend image from the Docker registry:
```bash
docker pull nnatchy14/backend:latest
```

---

## 3. Create a Custom Docker Network
Create a custom Docker network to enable communication between the backend and frontend containers:
```bash
docker network create my_network
```

---

## 4. Run Backend
Run the backend container and attach it to the same custom network:
```bash
docker run --network my_network -p 8010:8010 nnatchy14/backend:latest
```

---

## 5. Run Frontend
Run the frontend container and attach it to the same custom network:
```bash
docker run --network my_network -p 3000:3000 nnatchy14/frontend:latest
```

---

## How to Use This App

1. **Connect Your MetaMask Account**:
   - Visit the web application at `http://localhost:3000`.
   - Use your MetaMask wallet to connect your account. This step is necessary to register your identity securely.

2. **Register Your Information**:
   - Once connected, register your details (e.g., student ID, name).  
   - **Note**: Your student ID will be validated to ensure eligibility.

3. **Select the Voting Position**:
   - After registration, you will be able to choose the type of position you wish to vote for (e.g., class representative, club leader, etc.).

4. **Vote for Your Desired Candidate**:
   - Review the list of candidates for the selected position and cast your vote.

5. **View Results**:
   - After voting, you can view the results, including your vote and the overall distribution of votes from others.

---

### Notes:
1. **Communication Between Backend and Frontend**:
   - Ensure the backend and frontend are configured to communicate within the same network. For example, in the `.env` file or frontend configuration, reference the backend service using the container name (e.g., `http://backend:8010`).
   
2. **Access the Services**:
   - Access the backend via `http://localhost:8010`.
   - Access the frontend via `http://localhost:3000`.

3. **For More Information**:
   - If something isn’t working properly, check the respective `README.md` file in the `frontend` or `backend` directory for more details and troubleshooting steps.

4. **.env changes**
   - you must change `ETHEREUM_PRIVATE_KEY_VOTING=xxxx` with yours

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
