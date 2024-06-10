import React from 'react';
import { Carousel } from 'react-responsive-carousel';
import 'react-responsive-carousel/lib/styles/carousel.min.css';
import './home.css';

const AutoCarousel = ({ images }) => {
    return (
        <div className="carousel-container">
            <Carousel 
                autoPlay 
                interval={3000} 
                infiniteLoop 
                showThumbs={false} 
                showStatus={false}
                showArrows={true}
                stopOnHover={true}
            >
                {images.map((image, index) => (
                    <div key={index}>
                        <img src={image} alt={`Curso ${index + 1}`} />
                    </div>
                ))}
            </Carousel>
        </div>
    );
};

export default AutoCarousel;