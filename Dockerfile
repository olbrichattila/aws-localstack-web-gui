FROM php:8.3-apache

# Set the working directory to /var/www/html
WORKDIR /var/www/html

# Set the correct DocumentRoot for Apache
RUN sed -i -e 's!/var/www/html!/var/www/html/public!g' /etc/apache2/sites-available/000-default.conf

# Install dependencies
RUN apt-get update && \
    apt-get install -y \
        libzip-dev \
        zip \
        unzip

# Install PHP extensions
RUN docker-php-ext-install pdo_mysql zip

# Enable Apache modules
RUN a2enmod rewrite

# Copy composer files and install dependencies
COPY composer.json composer.lock ./
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer
RUN composer install --no-scripts --no-autoloader

# Install dependencies, including Node.js 21.6.0
RUN curl -fsSL https://deb.nodesource.com/setup_21.x | bash - && \
    apt-get install -y nodejs

# Copy the rest of the application code
COPY ./ /var/www/html

# Create required folders
RUN touch /var/www/html/database/database.sqlite \
    && mkdir -p /var/www/html/storage/app \
    && mkdir -p /var/www/html/storage/framework \
    && mkdir -p /var/www/html/storage/logs \
    && mkdir -p /var/www/html/storage/framework/sessions \
    && mkdir -p /var/www/html/storage/framework/views \
    && mkdir -p /var/www/html/bootstrap/cache \
    && chown www-data:www-data -R /var/www/html/storage \
    && chown www-data:www-data -R /var/www/html/bootstrap/cache \
    && chown www-data:www-data -R /var/www/html/storage/framework/views \
    && chown www-data:www-data -R /var/www/html//database

# Generate the Laravel autoload files
RUN composer dump-autoload --optimize \
    && php artisan migrate

RUN npm --prefix ./frontend/ install \
    && npm --prefix ./frontend/ run build

RUN cp -r ./frontend/build/static/ ./public/static/ \
    && cp ./frontend/build/index.html ./resources/views/index.blade.php

# Expose port 80
EXPOSE 80

# Start Apache
CMD ["apache2-foreground"]
