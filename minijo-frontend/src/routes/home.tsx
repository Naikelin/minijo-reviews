
import React from 'react';
import Container from '@mui/material/Container';

function Home() {
  return (
    <>
        <div className="bg-white">
      <div className="relative bg-gray-900">
        {/* Decorative image and overlay */}
        <div aria-hidden="true" className="absolute inset-0 overflow-hidden">
          <img
            src="https://www.cherrymx.de/_Resources/Persistent/1/d/e/9/1de92305b5e3b931c09daedcd7753fa21d57028c/DSC02907.jpg"
            alt=""
            className="w-full h-full object-center object-cover"
          />
        </div>
        <div
          aria-hidden="true"
          className="absolute inset-0 bg-gray-900 opacity-50"
        />

        <div className="relative max-w-3xl mx-auto py-32 px-6 flex flex-col items-center text-center sm:py-64 lg:px-0">
          <h1 className="text-4xl font-extrabold tracking-tight text-white lg:text-6xl">
            Teclados Mecánicos Customizados
          </h1>
          <p className="mt-4 text-xl text-white">
            Construye tu teclado paso a paso, eligiendo cada uno de los
            componentes
          </p>
          
            <p className="mt-8 inline-block bg-white border border-transparent rounded-md py-3 px-8 text-base font-medium text-gray-900 hover:bg-gray-100">
              Comienza tu personalizacion aqui!
            </p>
          
        </div>
      </div>
    </div>
    
    </>

  );
}

export default Home;