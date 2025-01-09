export default function Home() {
  return (
    <html lang="en">
      <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <meta http-equiv="X-UA-Compatible" content="ie=edge" />
        <title>JRWeinzapfel Training</title>
      </head>
      <body className="bg-gray-100 text-gray-900 font-sans">
        <header className="bg-white shadow-md">
          <div className="container mx-auto p-4">
            <img src="\images\training-logo.png"></img>
            <h1 className="text-4xl font-bold text-center">Home</h1>
            <nav className="mt-4">
              <ul className="flex justify-center space-x-4">
                <li>
                  <button className="text-lg font-medium text-gray-700 hover:text-gray-900">
                    Training
                  </button>
                </li>
                <li>
                  <button className="text-lg font-medium text-gray-700 hover:text-gray-900">
                    Programs
                  </button>
                </li>
                <li>
                  <button className="text-lg font-medium text-gray-700 hover:text-gray-900">
                    Blog
                  </button>
                </li>
                <li>
                  <button className="text-lg font-medium text-gray-700 hover:text-gray-900">
                    About
                  </button>
                </li>
                <li>
                  <button className="text-lg font-medium text-gray-700 hover:text-gray-900">
                    Contact
                  </button>
                </li>
              </ul>
            </nav>
          </div>
        </header>
        <main className="container mx-auto p-4">
          <article className="bg-white p-6 rounded-lg shadow-md">
            <p className="text-lg">Hello and welcome to my training page!</p>
          </article>
        </main>
      </body>
    </html>
  );
}
