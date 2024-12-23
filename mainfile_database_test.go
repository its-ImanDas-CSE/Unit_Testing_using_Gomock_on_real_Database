package main

import (
	"testing"

	"github.com/golang/mock/gomock"      // Gomock for mocking
	"github.com/stretchr/testify/assert" // Testify for assertions
	"gorm.io/gorm"
)

// TestGetStudentNameByID tests the GetStudentNameByID function using gomock.
func TestGetStudentNameByID(t *testing.T) {
	// Step 1: Create a gomock controller (this manages mock behavior)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // Ensure the controller checks all expected calls after the test.

	// Step 2: Create a mock version of the DBInterface
	mockDB := NewMockDBInterface(ctrl)

	// Step 3: Set expectations for the mock:
	// We expect the 'First' method to be called once with the ID 101,
	// and it should return a student with the name "Iman".
	mockDB.EXPECT().First(gomock.Any(), 101).DoAndReturn(func(out interface{}, where ...interface{}) *gorm.DB {
		// gomock.Any() means that the mock is not concerned with what out is, as long as it is some variable that can store the result.
		// 101 is the ID you are passing to the First() method, with variable.
		/*
			.DoAndReturn() is used to define a custom behavior when the First() method is called. It allows you to specify what should happen when First() is called in the test.
			 The function passed to .DoAndReturn() takes the following parameters:
			 1) out interface{}: This is the variable where the result (in our case, a Student) will be stored. In our case, out will be a pointer to a Student (which will be populated with mock data).
			 2) where ...interface{}: This is an optional parameter used to specify additional conditions. It can take multiple arguments, but we are only interested in the ID 101 in this case.

		*/
		student := Student{Name: "Iman"}
		// This line creates a new Student struct with a Name field set to "Iman".
		// This simulates the student data that What we want to return when the First() method is called.

		// Type assertion to ensure the data is stored correctly in 'out'.
		if s, ok := out.(*Student); ok {
			*s = student
		}
		// s will hold the value if the assertion is successful.
		//ok will be true if the type assertion succeeds and false if it fails.

		return &gorm.DB{} // Return an empty *gorm.DB instance.
	}).Times(1) // .Times(1) specifies that the First() method is expected to be called exactly once in the test.
	// in main file, the function that we kept under interface, that function can only be used in  ___.EXPECT().that_func....
	// because, mockgen create mock DB from interface.

	// Step 4: Call the function with the mock database and verify the result
	id := 101
	name, err := GetStudentNameByID(mockDB, id)

	// Step 5: Assert that the result is as expected
	assert.NoError(t, err)        // Ensure no error occurred
	assert.Equal(t, "Iman", name) // Ensure the name returned is "Iman"

}
