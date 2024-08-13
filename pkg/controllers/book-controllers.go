package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Muyi2905/GoLibManager/pkg/models"
	"github.com/Muyi2905/GoLibManager/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var Newbook models.Book

