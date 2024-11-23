import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import { fetchUserInfo, updateUserStatus } from "../libs/userLib"; // Ensure `updateUserStatus` is imported

// Async thunk for fetching user info
export const fetchUserInfoThunk = createAsyncThunk(
  "user/fetchUserInfo",
  async (walletAddress, { rejectWithValue }) => {
    try {
      console.log("in trunk:", walletAddress);
      const userInfo = await fetchUserInfo(walletAddress);
      return userInfo; // Return user info
    } catch (error) {
      return rejectWithValue(
        error.response?.data || "Failed to fetch user info"
      );
    }
  }
);

// Async thunk for updating user status
export const updateUserStatusThunk = createAsyncThunk(
  "user/updateUserStatus",
  async ({ walletAddress, updatePayload }, { rejectWithValue }) => {
    try {
      const response = await updateUserStatus(walletAddress, updatePayload);
      return {
        ...updatePayload, // Return the updated fields
        response, // Include the response in case needed later
      };
    } catch (error) {
      return rejectWithValue(
        error.response?.data || "Failed to update user status"
      );
    }
  }
);

const initialState = {
  userInfo: {
    id: null,
    walletAddress: null,
    username: null,
    isActive: false,
    studentId: null,
    firstName: null,
    lastName: null,
    hasVotedYear: false, // Use camelCase here
    hasVotedRepresentative: false, // Use camelCase here
  },
  loading: false,
  error: null,
};

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    resetUserInfo: () => initialState, // Reset state to initial
  },
  extraReducers: (builder) => {
    builder
      // Handle fetchUserInfoThunk
      .addCase(fetchUserInfoThunk.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchUserInfoThunk.fulfilled, (state, action) => {
        state.loading = false;
        // Map the API response to camelCase fields
        state.userInfo = {
          ...action.payload,
          walletAddress: action.payload.wallet_address, // Map wallet_address to walletAddress
          hasVotedYear: action.payload.has_vote_year,
          hasVotedRepresentative: action.payload.has_vote_rep,
        };
      })
      .addCase(fetchUserInfoThunk.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      })

      // Handle updateUserStatusThunk
      .addCase(updateUserStatusThunk.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(updateUserStatusThunk.fulfilled, (state, action) => {
        state.loading = false;
        // Update the specific fields in the userInfo state
        state.userInfo.hasVotedYear =
          action.payload.has_vote_year ?? state.userInfo.hasVotedYear;
        state.userInfo.hasVotedRepresentative =
          action.payload.has_vote_rep ?? state.userInfo.hasVotedRepresentative;
      })
      .addCase(updateUserStatusThunk.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      });
  },
});

export const { resetUserInfo } = userSlice.actions;
export default userSlice.reducer;
