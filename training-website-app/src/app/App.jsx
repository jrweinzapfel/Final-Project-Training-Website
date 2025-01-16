import Navbar from './components/Navbar';
import { About, Contact, Programs, Training, Blog } from '@/components/pages';
import Home from './page';


function App() {
    return (
        <div>
            <Navbar />
            <Routes>
                <Route path="/" element={<Home />}/>
                <Route path="/training" element={<Training />} />
                <Route path="/programs" element={<Programs />} />
                <Route path="/blog" element={<Blog />} />
                <Route path="/about" element={<About />}/>
                <Route path="/contact" element={<Contact />}/>
            </Routes>
        </div>
    )
}