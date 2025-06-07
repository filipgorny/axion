package example

import (
	"github.com/filipgorny/axion/pkg/node"
	"github.com/filipgorny/axion/pkg/node/folder"
	"github.com/filipgorny/axion/pkg/node/model"
	"github.com/filipgorny/axion/pkg/node/service"
	"github.com/filipgorny/axion/pkg/project"
)

func NewExample() *project.Project {
	rootFolder := folder.NewFolderNode("root")

	var project = &project.Project{
		Name:     "Example Project",
		RootNode: rootFolder,
	}

	rootFolder.AddElement(
		folder.NewFolderNode("src"),
		folder.NewFolderNode("entities"),
		folder.NewFolderNode("components"),
		folder.NewFolderNode("dtos"),
		folder.NewFolderNode("services"),
		folder.NewFolderNode("utils"),
		folder.NewFolderNode("views"),
		folder.NewFolderNode("assets"),
		folder.NewFolderNode("styles"),
		folder.NewFolderNode("tests"),
		folder.NewFolderNode("config"),
	)

	userEntity := model.NewEntity("User").AddProperty("UserID", node.ID).AddProperty("UserName", node.String).AddProperty("UserEmail", node.String)

	rootFolder.FindChildFolder(node.NewNodeQuery("entities")).AddElement(
		userEntity,
		model.NewEntity("Product").AddProperty("ProductID", node.ID).AddProperty("ProductName", node.String).AddProperty("ProductPrice", node.Money),
		model.NewEntity("Order").AddProperty("OrderID", node.ID).AddProperty("OrderDate", node.DateTime).AddProperty("OrderTotal", node.Money),
	)

	serviceInstance := service.NewServiceNode("UserService")

	serviceInstance.AddMethods(
		service.NewMethod("GetUserByID", service.NewReturn("user", node.Entity), service.NewParameter("userID", node.ID), service.NewParameter("includeDetails", node.Boolean)),
		service.NewMethod("CreateUser", service.NewVoidReturn(), service.NewEntityParameter("user", userEntity)),
		service.NewMethod("UpdateUser", service.NewVoidReturn(), service.NewEntityParameter("user", userEntity)),
		service.NewMethod("DeleteUser", service.NewVoidReturn(), service.NewEntityParameter("user", userEntity)),
	)

	rootFolder.FindChildFolder(node.NewNodeQuery("services")).AddElement(serviceInstance)

	return project
}
