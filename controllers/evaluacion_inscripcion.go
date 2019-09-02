package controllers

import (
	//"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/planesticud/campus_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// AdmisionController ...
type EvaluacionInscripcionController struct {
	beego.Controller
}

// URLMapping ...
func (c *EvaluacionInscripcionController) URLMapping() {
	c.Mapping("PutEvaluacionInscripcion", c.PutEvaluacionInscripcion)
	c.Mapping("GetEvaluacionInscripcion", c.GetEvaluacionInscripcion)
	c.Mapping("GetEvaluacionInscripcionByIdInscripcion", c.GetEvaluacionInscripcionByIdInscripcion)
}


// PutEvaluacionInscripcion ...
// @Title PutEvaluacionInscripcionAdmision
// @Description Modificar notas de la evaluación
// @Param	id		path 	string	true		"el id de la evaluacion a modificar"
// @Param	body		body 	{}	true		"body Modificar Evaluacion content"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /:id [put]
func (c *EvaluacionInscripcionController) PutEvaluacionInscripcion() {

}

// GetEvaluacionInscripcion ...
// @Title GetEvaluacionInscripcion
// @Description consultar evaluacion por inscripcion_id
// @Param	inscripcion_id		path 	string	true		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EvaluacionInscripcionController) GetEvaluacionInscripcionByIdInscripcion() {
	//Id de la inscripcion a consultar
	idStr := c.Ctx.Input.Param(":id")
	fmt.Println("Id de inscripcion: ", idStr)
 	// alerta que retorna la funcion GetEvaluacionInscripcionByIdInscripcion
	var alerta models.Alert
	//cadena de alertas
  alertas := append([]interface{}{"Cadena de respuestas: "})
	// //cadena de alertas
	// alertas := append([]interface{}{"Cadena de respuestas: "})
	//resultado admision
	var resultado []map[string]interface{}
	var resultadoEntrevista []map[string]interface{}
	var entrevista map[string]interface{}
	var resultado2 map[string]interface{}
	var resultado3 map[string]interface{}
	// var puntajeTotal map[string]interface{}
	// fmt.Println("EL API es: ", beego.AppConfig.String("EvaluacionInscripcionService"))

	// Obtener notas por InscripcionId
	errEvaluacionInscripcion := request.GetJson("http://"+beego.AppConfig.String("EvaluacionInscripcionService")+"/evaluacion_inscripcion/?query=InscripcionId:"+idStr, &resultado)
	if errEvaluacionInscripcion == nil && resultado != nil {
		if resultado[0]["Type"] != "error" {
			var notaAcumulada = 0.00
			var notaEntrevista = 0.00
			var entrevistaId = ""
			var evaluacionId = ""
			var entrevistadores = 0
			for u := 0; u < len(resultado); u++ {
				// fmt.Println("Nota: ", resultado[u]["NotaFinal"])
				var requisitoProgramaAcademico map[string]interface{}
				requisitoProgramaAcademico = resultado[u]["RequisitoProgramaAcademicoId"].(map[string]interface{})
				//Calcular nota si el criterio tiene entrevista
				if resultado[u]["EntrevistaId"] != nil {
					entrevista = resultado[u]["EntrevistaId"].(map[string]interface{})
					entrevistaId = fmt.Sprintf("%.f", entrevista["Id"].(float64))
					evaluacionId = fmt.Sprintf("%.f", resultado[u]["Id"])  //new
					// fmt.Println("Id de la entrevista", entrevista["Id"])
					errEntrevistadorEntrevista := request.GetJson("http://"+beego.AppConfig.String("EvaluacionInscripcionService")+"/entrevistador_entrevista/?query=EntrevistaId:"+entrevistaId+"&EstadoEntrevistaId:3", &resultadoEntrevista)
					if errEntrevistadorEntrevista == nil && resultadoEntrevista != nil {
						for u := 0; u< len(resultadoEntrevista); u++ {
								notaEntrevista += resultadoEntrevista[u]["NotaParcial"].(float64)
						}
						entrevistadores = len(resultadoEntrevista)
						notaEntrevista = notaEntrevista/float64(entrevistadores)
						resultado[u]["NotaFinal"] = notaEntrevista //New
						var evaluacionUpdate map[string]interface{}
						// New Actualiza el valor de la nota de la entrevista en el registro
						evaluacionUpdate = resultado[u]
						errSolicitud := request.SendJson("http://"+beego.AppConfig.String("EvaluacionInscripcionService")+"/evaluacion_inscripcion/"+evaluacionId, "PUT", &resultado2, evaluacionUpdate)
						if errSolicitud == nil {
	            if resultado2["Type"] == "success" {
								alertas = append(alertas, "OK UPDATE evaluacion_inscripcion")
								alerta.Code = "200"
								alerta.Type = "success"
								alerta.Body = alertas
							}
	          } else {
	  					alertas = append(alertas, errSolicitud.Error())
	  					alerta.Code = "400"
	  					alerta.Type = "error"
						}
						///  Here new

						notaAcumulada += notaEntrevista * requisitoProgramaAcademico["Porcentaje"].(float64) / 100.00
					}
				} else {
					notaAcumulada += resultado[u]["NotaFinal"].(float64) * requisitoProgramaAcademico["Porcentaje"].(float64) / 100.00
				}
				resultado[u]["NotaFinal"] = notaEntrevista
				var resultadoInscripcion map[string]interface{}
				errInscripcion := request.GetJson("http://"+beego.AppConfig.String("AdmisionService")+"/inscripcion/"+idStr, &resultadoInscripcion)
				if errInscripcion == nil && resultadoInscripcion != nil {
					if resultadoInscripcion["Type"] != "error" {
						resultadoInscripcion["PuntajeTotal"] = notaAcumulada
						errSolicitud := request.SendJson("http://"+beego.AppConfig.String("AdmisionService")+"/inscripcion/"+fmt.Sprintf("%.f", resultadoInscripcion["Id"].(float64)), "PUT", &resultado3, resultadoInscripcion)
						if errSolicitud == nil {
							if resultado3["Type"] == "success" {
								alertas = append(alertas, "OK UPDATE inscripcion")
								alerta.Code = "200"
								alerta.Type = "success"
								alerta.Body = alertas
							}
							// fmt.Println("Actualizado CORRECTAMENTE", resultadoInscripcion)
						} else {
							alertas = append(alertas, errSolicitud.Error())
							alerta.Code = "400"
							alerta.Type = "error"
							// fmt.Println("Codigo de ERROR: ", errSolicitud)
						}
					}
				} else {
						fmt.Println("ERROR: " + errInscripcion.Error())
				}
				alerta.Body = notaAcumulada
			}
		}
	} else {
		fmt.Println ("Error")
		if errEvaluacionInscripcion != nil {
			c.Data["json"] = nil
		} else {
			c.Data["json"] = nil
		}
	}
	c.Data["json"] = alerta
	c.ServeJSON()
}


//calcular nota final de la inscripcion
func (c *EvaluacionInscripcionController) GetEvaluacionInscripcion() {
	idStr := c.Ctx.Input.Param(":id")
	fmt.Println("Id Inscripcion es: ", idStr)
	// //alerta que retorna la funcion GetEvaluacionInscripcion
	var alerta models.Alert
	// var alerta models.Alert
	// alertas := append([]interface{}{"Cadena de respuestas: "})
	var resultado []map[string]interface{}
	errEvaluacionInscripcion := request.GetJson("http://"+beego.AppConfig.String("EvaluacionInscripcionService")+"/evaluacion_inscripcion/?query=InscripcionId:"+idStr, &resultado)
		if errEvaluacionInscripcion == nil && resultado != nil {
			if resultado[0]["Type"] != "error" {
				//consultar porcentaje de requisito
				var notaAcumulada = 0.00
				for u := 0; u < len(resultado); u++ {
					var requisitoProgramaAcademico map[string]interface{}
					requisitoProgramaAcademico = resultado[u]["RequisitoProgramaAcademicoId"].(map[string]interface{})
					notaAcumulada += resultado[u]["NotaFinal"].(float64) * requisitoProgramaAcademico["Porcentaje"].(float64) / 100.00
					alerta.Body = notaAcumulada
				}
			}
		} else {
			fmt.Println ("Error")
				if errEvaluacionInscripcion != nil {
					c.Data["json"] = nil
				} else {
					c.Data["json"] = nil
				}
		}
		c.Data["json"] = alerta
		c.ServeJSON()
}
