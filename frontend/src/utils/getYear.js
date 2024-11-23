export const getYearFromStudentId = (studentId) => {
    if (!studentId || studentId.length < 2) {
      throw new Error("Invalid student ID");
    }
  
    const yearPrefix = parseInt(studentId.slice(0, 2), 10);
    const currentYear = (new Date().getFullYear() + 543) % 100;
  
    if (yearPrefix === currentYear) return 1;
    if (yearPrefix === currentYear - 1) return 2;
    if (yearPrefix === currentYear - 2) return 3;
    if (yearPrefix === currentYear - 3) return 4;
  
    throw new Error("Year could not be determined from the student ID");
  };
  