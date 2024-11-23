
# adv-top-proj

```bash
By Group Mairoira:
- 6431307421 Jirawat Lengnoi
- 6431318321 Natthasith Wiriyayothin
- 6431342321 Wish Marukapitak
```

---

## Docker Image Setup Instructions

Follow these steps to set up and run the backend and frontend Docker containers, ensuring they communicate effectively.

---

### 1. Pull Backend Image
Download the backend image from the Docker registry:
```bash
docker pull nnatchy14/backend:latest
```

---

### 2. Pull Frontend Image
Download the frontend image from the Docker registry:
```bash
docker pull nnatchy14/frontend:latest
```

---

### 3. Create a Custom Docker Network
Create a custom Docker network for communication between the backend and frontend containers:
```bash
docker network create my_network
```

---

### 4. Run Backend
Run the backend container, attaching it to the custom network:
```bash
docker run --network my_network -p 8010:8010 nnatchy14/backend:latest
```

---

### 5. Run Frontend
Run the frontend container, attaching it to the custom network:
```bash
docker run --network my_network -p 3000:3000 nnatchy14/frontend:latest
```

---

## How to Use the Application

1. **Connect Your MetaMask Account**:
   - Visit the web application at `http://localhost:3000`.
   - Use your MetaMask wallet to securely connect your account.

2. **Register Your Information**:
   - After connecting your wallet, register your details (e.g., student ID, name).
   - **Note**: Your student ID will be validated to ensure eligibility.

3. **Select the Voting Position**:
   - Choose the position for which you want to vote (e.g., class representative, club leader).

4. **Vote for Your Desired Candidate**:
   - Review the candidates for the selected position and cast your vote.

5. **View Results**:
   - After voting, check the results, including your vote and the overall distribution of votes.

---

### Additional Notes

1. **Backend-Frontend Communication**:
   - Ensure both services are configured to communicate within the same Docker network. Use the container name (e.g., `http://backend:8010`) in the frontend or backend configurations.

2. **Service Access**:
   - Access the backend via `http://localhost:8010`.
   - Access the frontend via `http://localhost:3000`.

3. **Environment Variables**:
   - Update the `.env` file, particularly the `ETHEREUM_PRIVATE_KEY_VOTING` variable:
     ```env
     ETHEREUM_PRIVATE_KEY_VOTING=your_private_key
     ```
     **Note**: The `ETHEREUM_PRIVATE_KEY_VOTING` must be the private key of the wallet address. This is required because the team cannot find a way to include sensitive environment variables in the Docker images.

4. **GitLab References**:
   - Initial Frontend Repository: [Frontend GitLab Repository](https://gitlab.com/adv-top-project/frontend)
   - Initial Backend Repository: [Backend GitLab Repository](https://gitlab.com/adv-top-project/backend)

5. **For More Information**:
   - Refer to the `README.md` files in the `frontend` or `backend` directories for detailed instructions or troubleshooting.

---

### Troubleshooting

- **Network Issues**:
  - Verify that both containers are in the same Docker network:
    ```bash
    docker network inspect my_network
    ```

- **Environment Variables**:
  - Double-check your `.env` file for correct backend and blockchain configurations.

--- 
