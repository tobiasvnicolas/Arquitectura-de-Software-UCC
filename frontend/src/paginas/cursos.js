import React from 'react'
import   './cursos.css'

const Cursos = () => {
    return (
      //La forma de poder darle stilos a las pag en react es a traves de className={styles.}, en este caso le estasmos dando el estilo al fondo de la misma
      <div className="fondo"> <br />
    
            {/*La forma de encuadrar un contenido es a través del bloque de construcción container*/}
            <div class="container">
    
                <div class="card text-bg-dark">
                    
                    {/*En react las imagenes se agregan a través del codigo src={require("...")}*/}
        
                    <div class="card-img-overlay">
        
                    <h1 class="card-title">Bienvenido a la Sección Cursos</h1>
                    <h3 class="card-text">En este apartado te mostraremos todos los cursos con los que contamos.</h3>
        
                    </div>
        
                </div>
        
            </div>
    
    
            <br/><br/><br/>
    
            {/*En este caso el className={styles.tarjetas} se encargar de darle el tamaño necesario a nuestras tarjetas*/}
            <div className="tarjetas">
    
            {/*En className={styles.card} lo que hacemos es darle ese efecto de zoom a la hora de pasar el cursor por encima de la tarjeta*/}
            <div className="card">
    
                <div class="card mb-4 bg-secondary-subtle">
    
                <div class="row g-0">
                    
                    {/*Con col-md-2 modificamos el tamaño de la tarjeta*/}
                    <div class="col-md-2">
    
    
                    </div>
    
                    <div class="col-md-8">
    
                    <div class="card-body">
    
                        <h5 class="card-title"><a href={"/cursos/" + 1}>HTML desde 0</a></h5>
    
                        <p class="card-text">Nuestro programa está diseñado para equiparte con habilidades avanzadas en desarrollo backend utilizando JAVA,
                        incluyendo frameworks como Spring Boot, manejo de bases de datos, y mejores prácticas en seguridad y diseño de software.</p>
                        
                        <p class="card-text"><small class="text-body-secondary">
                        • <b>Requisitos:</b> <br/>
                        &nbsp;&nbsp;- Ninguno. <br/>
                        • <b>Duración:</b> <br/>
                        &nbsp;&nbsp;- 4 Semanas. <br />
                        • <b>Profesor:</b> <br/>
                        &nbsp;&nbsp;- Juan Pérez<br />
                        • <b>Categoria:</b> <br/>
                        &nbsp;&nbsp;- HTML.
                        </small></p>
    
                    </div>
    
                    </div>
    
                </div>
    
                </div>
    
            </div>
            
            <div className="card">
    
                <div class="card mb-4 bg-secondary-subtle">
    
                <div class="row g-0">
    
                    <div class="col-md-2">
    
    
                    </div>
    
                    <div class="col-md-8">
    
                    <div class="card-body">
    
                        <h5 class="card-title">JavaScript</h5>
    
                        <p class="card-text">Aprende a construir sitios web desde cero con nuestro curso introductorio de HTML.
                             Este curso te guiará a través de los fundamentos esenciales de HTML, el lenguaje estándar para la creación de páginas web. 
                            Desde la estructura básica de un documento HTML hasta la creación de formularios interactivos y la integración de imágenes y enlaces, dominarás las herramientas necesarias para crear sitios web atractivos y funcionales.</p>
                        
                        {/*EL text-body-secondary se utiliza para darle el color de fondo a las tarjetas*/}
                        <p class="card-text"><small class="text-body-seconadary">
                        • <b>Requisitos:</b> <br/>
                        &nbsp;&nbsp;- Conocimientos Previos: No se requieren conocimientos previos de HTML ni de desarrollo web. El curso está diseñado para principiantes absolutos. <br/>
                        &nbsp;&nbsp;- Herramientas: Es recomendable tener acceso a un ordenador con conexión a Internet para practicar y seguir las lecciones interactivas.<br/>
                        &nbsp;&nbsp;&nbsp;- Capacidad de Autodidactismo: Aunque no se requiere experiencia previa, es beneficioso tener habilidades básicas de autodidactismo y capacidad para seguir instrucciones. <br/>
                        • <b>Duración:</b> <br/>
                        &nbsp;&nbsp;- 10 Semanas<br />
                        • <b>Profesor:</b> <br/>
                        &nbsp;&nbsp;- María Gómez<br />
                        • <b>Categoria:</b> <br/>
                        &nbsp;&nbsp;- Programación.
                        </small></p>
    
                    </div>
    
                    </div>
    
                </div>
    
                </div>
    
            </div>
    
            <div className="card">
    
                <div class="card mb-4 bg-secondary-subtle">
    
                <div class="row g-0">
    
                    <div class="col-md-2">
    
                    
                    </div>
    
                    <div class="col-md-8">
    
                    <div class="card-body">
    
                        <h5 class="card-title">C++</h5>
    
                        <p class="card-text">El curso de C++ está diseñado para proporcionar a los estudiantes una comprensión profunda y práctica de uno de los lenguajes de programación más potentes y versátiles. 
                            Desde los fundamentos básicos hasta técnicas avanzadas, los participantes explorarán cómo utilizar C++ para desarrollar software eficiente y robusto. 
                            A lo largo del curso, se abordarán conceptos clave como la programación orientada a objetos, el manejo de memoria, la sintaxis avanzada y las mejores prácticas de desarrollo de software..</p>
    
                        <p class="card-text"><small class="text-body-secondary">
                        • <b>Requisitos:</b> <br/>
                        &nbsp;&nbsp;-Conocimientos Previos: No se requieren conocimientos previos de C++ o programación, aunque tener una comprensión básica de la lógica de programación puede ser ventajoso. <br/>
                        &nbsp;&nbsp;- Herramientas: Se necesita acceso a un ordenador con un compilador de C++ instalado (como GCC, Visual Studio o Clang) y un entorno de desarrollo integrado (IDE) como Code::Blocks, Visual Studio Code o Eclipse. <br/>
                        &nbsp;&nbsp;- Motivación y Compromiso: Es fundamental tener la disposición para aprender y dedicar tiempo a completar las tareas y ejercicios prácticos del curso. <br/>
                        &nbsp;&nbsp;- Capacidad de Resolución de Problemas: Habilidad para resolver problemas de manera lógica y sistemática, así como la capacidad de analizar y entender conceptos abstractos. <br/>
                        • <b>Duración:</b> <br/>
                        &nbsp;&nbsp;- 8 Semanas.<br />
                        • <b>Profesor:</b> <br/>
                        &nbsp;&nbsp;- Carlos Sánchez.<br />
                        • <b>Categoria:</b> <br/>
                        &nbsp;&nbsp;- Programación.
                        </small></p>
    
                    </div>
    
                    </div>
    
                </div>
    
                </div>
    
            </div>
            
            <div className="card">
    
                <div class="card mb-4 bg-secondary-subtle">
    
                <div class="row g-0">
    
                    <div class="col-md-2">
    
    
                    </div>
    
                    <div class="col-md-8">
    
                    <div class="card-body">
    
                        <h5 class="card-title">Golang</h5>
    
                        <p class="card-text">El curso de Go (Golang) está diseñado para introducir a los estudiantes en el lenguaje de programación Go, conocido por su simplicidad, eficiencia y facilidad para construir software robusto. 
                            A lo largo del curso, los participantes explorarán desde los conceptos básicos hasta técnicas avanzadas de programación en Go, incluyendo concurrencia y manejo de errores. 
                            Este curso es ideal tanto para principiantes como para desarrolladores que desean ampliar sus habilidades en un lenguaje de programación moderno y en crecimiento.</p>
                        
                        <p class="card-text"><small class="text-body-secondary">
                        • <b>Requisitos:</b> <br/>
                        &nbsp;&nbsp;- Conocimientos Previos: No se requieren conocimientos previos de Go, pero tener experiencia básica en programación y comprensión de conceptos como variables, tipos de datos y estructuras de control será beneficioso. <br/>
                        &nbsp;&nbsp;- Herramientas: Acceso a un ordenador con un compilador de Go instalado y un entorno de desarrollo integrado (IDE) como VSCode, GoLand o Atom. <br/>
                        &nbsp;&nbsp;- Motivación y Compromiso: Es esencial estar motivado para aprender y dedicar tiempo a completar las tareas y ejercicios del curso.<br/>
                        &nbsp;&nbsp;- Capacidad de Aprendizaje: Habilidad para asimilar conceptos abstractos y resolver problemas de manera lógica y estructurada. <br/>
                        • <b>Duración:</b> <br/>
                        &nbsp;&nbsp;- 5 semanas<br />
                        • <b>Profesor:</b> <br/>
                        &nbsp;&nbsp;- Ana López<br />
                        • <b>Categoria:</b> <br/>
                        &nbsp;&nbsp;- Programación.
                        </small></p>
    
                    </div>
                    </div>
    
                </div>
    
                </div>
    
            </div>
    
            <div className="card">
    
                <div class="card mb-4 bg-secondary-subtle">
    
                <div class="row g-0">
    
                    <div class="col-md-2">
    
    
                    </div>
    
                    <div class="col-md-8">
    
                    <div class="card-body">
    
                        <h5 class="card-title">Python</h5>
    
                        <p class="card-text">El curso de Python es una introducción completa al lenguaje de programación Python, diseñado para principiantes y aquellos que desean fortalecer sus fundamentos en programación. 
                            A lo largo del curso, los estudiantes aprenderán desde los conceptos básicos hasta técnicas avanzadas de programación en Python. 
                            Se explorarán aplicaciones prácticas y casos de uso real, abarcando desde el desarrollo de scripts hasta la creación de aplicaciones web utilizando frameworks modernos.</p>
                        
                        <p class="card-text"><small class="text-body-secondary">
                        • <b>Requisitos:</b> <br/>
                        &nbsp;&nbsp;- Conocimientos Previos: No se requieren conocimientos previos de Python. Sin embargo, tener una comprensión básica de la lógica de programación será útil. <br/>
                        &nbsp;&nbsp;- Herramientas: Acceso a un ordenador con Python instalado (preferiblemente Python 3.x) y un entorno de desarrollo integrado (IDE) como VSCode, PyCharm o Jupyter Notebook. <br/>
                        &nbsp;&nbsp;- Motivación y Compromiso: Es esencial estar motivado para aprender y dedicar tiempo a completar las tareas y ejercicios del curso. <br/>
                        &nbsp;&nbsp;- Capacidad de Aprendizaje: Habilidad para asimilar conceptos abstractos y resolver problemas de manera lógica y estructurada. <br/>
                        • <b>Duración:</b> <br/>
                        &nbsp;&nbsp;- 8 Semanas.<br />
                        • <b>Profesor:</b> <br/>
                        &nbsp;&nbsp;- Profesor D.<br />
                        • <b>Categoria:</b> <br/>
                        &nbsp;&nbsp;- Programación.
                        </small></p>
    
                    </div>
    
                    </div>
    
                </div>
    
                </div>
    
            </div>
    
    
        
    
            </div>
    
            <div className="piePagina">
    
            <fieldset>
            
                <legend>Sobre el Curso</legend>
                
                <p>Udemy es una <b>plataforma de educación online en vivo</b> creada con la misión de democratizar la educacion de calidad en toda Latinoamérica. Queremos
                que nuestros alumnos estén preparados de la mejor manera para su salidad al mundo laboral.</p>
                
            </fieldset>
    
            </div>
    
      </div>
  
    )
  
  }
  
  export default Cursos