package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:AdmisionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:AdmisionController"],
        beego.ControllerComments{
            Method: "PutAdmision",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:AdmisionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:AdmisionController"],
        beego.ControllerComments{
            Method: "GetAdmision",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"],
        beego.ControllerComments{
            Method: "PostDescuentoAcademico",
            Router: `/descuentoacademico`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"],
        beego.ControllerComments{
            Method: "GetDescuentoAcademico",
            Router: `/descuentoacademico/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"],
        beego.ControllerComments{
            Method: "PutDescuentoAcademico",
            Router: `/descuentoacademico/:Id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"],
        beego.ControllerComments{
            Method: "DeleteDescuentoAcademico",
            Router: `/descuentoacademico/:Id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"],
        beego.ControllerComments{
            Method: "GetDescuentoAcademicoByPersona",
            Router: `/descuentoacademicopersona/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"],
        beego.ControllerComments{
            Method: "GetDescuentoByDependenciaPeriodo",
            Router: `/descuentodependenciaperiodo/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:DescuentoController"],
        beego.ControllerComments{
            Method: "GetDescuentoByPersonaPeriodoDependencia",
            Router: `/descuentopersonaperiododependencia/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:EvaluacionInscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:EvaluacionInscripcionController"],
        beego.ControllerComments{
            Method: "PutEvaluacionInscripcion",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:EvaluacionInscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:EvaluacionInscripcionController"],
        beego.ControllerComments{
            Method: "GetEvaluacionInscripcionByIdInscripcion",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "PostExperienciaLaboral",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "PutExperienciaLaboral",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "GetExperienciaLaboral",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ExperienciaLaboralController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ExperienciaLaboralController"],
        beego.ControllerComments{
            Method: "DeleteExperienciaLaboral",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:FormacionController"],
        beego.ControllerComments{
            Method: "PostFormacionAcademica",
            Router: `/formacionacademica`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:FormacionController"],
        beego.ControllerComments{
            Method: "GetFormacionAcademica",
            Router: `/formacionacademica/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:FormacionController"],
        beego.ControllerComments{
            Method: "PutFormacionAcademica",
            Router: `/formacionacademica/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:FormacionController"],
        beego.ControllerComments{
            Method: "DeleteFormacionAcademica",
            Router: `/formacionacademica/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:FormacionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:FormacionController"],
        beego.ControllerComments{
            Method: "GetFormacionAcademicaByEnte",
            Router: `/formacionacademicaente/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:OrganizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:OrganizacionController"],
        beego.ControllerComments{
            Method: "GetByIdentificacion",
            Router: `/identificacion/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:OrganizacionController"],
        beego.ControllerComments{
            Method: "GetByEnte",
            Router: `/identificacionente/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "ActualizarPersona",
            Router: `/ActualizarPersona`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "ConsultaPersona",
            Router: `/ConsultaPersona/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "DatosComplementariosPersona",
            Router: `/DatosComplementarios`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "ActualizarDatosComplementarios",
            Router: `/DatosComplementarios`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "ConsultaDatosComplementarios",
            Router: `/DatosComplementarios/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "GuardarDatosContacto",
            Router: `/DatosContacto`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "ActualizarDatosContacto",
            Router: `/DatosContacto`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "DatosContacto",
            Router: `/DatosContacto/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "GuardarPersona",
            Router: `/GuardarPersona`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "RegistrarUbicaciones",
            Router: `/RegistrarUbicaciones`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "PostProduccionAcademica",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "PutProduccionAcademica",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "DeleteProduccionAcademica",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "GetProduccionAcademica",
            Router: `/:persona`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ProduccionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/campus_mid/controllers:ProduccionAcademicaController"],
        beego.ControllerComments{
            Method: "PutEstadoAutorProduccionAcademica",
            Router: `/estado_autor_produccion/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
