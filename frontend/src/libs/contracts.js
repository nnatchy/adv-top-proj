import axiosInstance from '../stores/axiosInstance';

export const setNumberInSimpleStorage = async (value) => {
  try {
    const response = await axiosInstance.post('/storage/v1/set', { value });
    return response.data;
  } catch (error) {
    console.error('Error setting number in SimpleStorage:', error);
    throw error;
  }
};

export const getNumberFromSimpleStorage = async () => {
  try {
    const response = await axiosInstance.get('/storage/v1/get');
    return response.data;
  } catch (error) {
    console.error('Error getting number from SimpleStorage:', error);
    throw error;
  }
};
