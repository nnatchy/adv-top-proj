# Frontend: NFT Authentication and Ethereum Voting System

This is the **frontend service** for the NFT-based authentication and Ethereum-powered voting platform. It is built using **Next.js** with **TypeScript** and styled with **Tailwind CSS**. The application runs on port `3000`.

---

## Features

- **User Authentication**:
  - Secure login via NFT-based authentication.
  - Seamless integration with Ethereum smart contracts.

- **Voting Interface**:
  - User-friendly UI for casting and tracking votes.
  - Real-time updates powered by blockchain interactions.

- **Responsive Design**:
  - Fully responsive design optimized for desktop and mobile devices.

---

## Prerequisites

- **Node.js** (version 16 or higher)
- **npm**, **yarn**, or **pnpm** package manager

---

## Getting Started

Follow these steps to set up and run the project:

### 1. Clone the Repository
```bash
git clone https://gitlab.com/adv-top-project/frontend.git
cd frontend
```

### 2. Install Dependencies
Run the following command to install all required Node.js modules:
```bash
npm install
# or
yarn install
```

### 3. Run the Development Server
```bash
npm run dev
# or
yarn dev
```

The development server will be available at:

```arduino
http://localhost:3000
```

---

## Project Structure
```bash
├── public                 # Static assets (e.g., images, fonts)
├── src                    # Source code
│   ├── app                # Application entry and routing
│   ├── components         # Reusable UI components
│   ├── services           # API service logic
│   └── styles             # Global and component-specific styles
├── package.json           # Project dependencies and scripts
├── tailwind.config.js     # Tailwind CSS configuration
├── tsconfig.json          # TypeScript configuration
└── README.md              # Project documentation
```

---

## Technologies Used
  - **Next.js**: React framework for server-rendered applications.
  - **TypeScript**: Enhances code reliability and maintainability.
  - **Tailwind CSS**: Utility-first CSS framework for rapid UI development.
  - **Ethereum**: Blockchain platform integrated with frontend for voting and NFT operations.

---

## Deployment

You can deploy the application to any modern hosting platform. For optimal performance and simplicity, consider using [Vercel](https://vercel.com/), the creators of Next.js. 

### Steps:
1. Deploy the repository directly on Vercel.
2. Configure environment variables under "Settings" -> "Environment Variables".

---

## Contributing

We welcome contributions! To get started:
1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/your-feature
   ```
3. Commit your changes and open a pull request.

---
