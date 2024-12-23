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
	// and it should return a student with the name "John Doe".
	mockDB.EXPECT().First(gomock.Any(), 101).DoAndReturn(func(out interface{}, where ...interface{}) *gorm.DB {
		student := Student{Name: "Iman"}
		// Type assertion to ensure the data is stored correctly in 'out'
		if s, ok := out.(*Student); ok {
			*s = student
		}
		return &gorm.DB{} // Return an empty *gorm.DB instance
	}).Times(1)

	// Step 4: Call the function with the mock database and verify the result
	id := 101
	name, err := GetStudentNameByID(mockDB, id)

	// Step 5: Assert that the result is as expected
	assert.NoError(t, err)        // Ensure no error occurred
	assert.Equal(t, "Iman", name) // Ensure the name returned is "John Doe"
}
