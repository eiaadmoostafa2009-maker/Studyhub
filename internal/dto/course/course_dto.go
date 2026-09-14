package course

type (
	//upload file
	UploadFileRequest struct{
        FileName  string `form:"file_name" validate:"required"`
	}
	UploadFileResponse struct{
		ID int
	}
	//share
	ShareCourseRequest struct {
		FileID      int `form:"file_id" validate:"required"`
		Title          string `form:"title" validate:"required"`
	}
	
	//update
	UpdateCourseRequest struct {
		Title          string `form:"title" validate:"required"`
		FileID     int `form:"file_id" validate:"required"`
	}
	
	//delete
	DeleteCourseResponse struct {
		Message string
	}
	//show all

	ShowCoursesResponse struct {
	    ID        int
		UserName string
	    Title     string
	    FileURL  string
    }  
)
