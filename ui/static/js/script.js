
        document.addEventListener('DOMContentLoaded', () => {
            const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);

            if ($navbarBurgers.length > 0) {
                $navbarBurgers.forEach(el => {
                    el.addEventListener('click', () => {
                        const target = el.dataset.target;
                        const $target = document.getElementById(target);

                        el.classList.toggle('is-active');
                        $target.classList.toggle('is-active');
                    });
                });
            }

            // Animate elements when they come into view
            const animateOnScroll = () => {
                const elements = document.querySelectorAll('.animate__animated');

                elements.forEach(element => {
                    const elementPosition = element.getBoundingClientRect().top;
                    const windowHeight = window.innerHeight;

                    if (elementPosition < windowHeight - 100) {
                        const animationClass = element.classList.contains('animate__fadeIn') ? 'animate__fadeIn' :
                                              element.classList.contains('animate__fadeInUp') ? 'animate__fadeInUp' : '';
                        if (animationClass) {
                            element.style.opacity = '1';
                        }
                    }
                });
            };

            window.addEventListener('scroll', animateOnScroll);
            animateOnScroll(); // Run once on page load
        });