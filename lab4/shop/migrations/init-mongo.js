// Initialize the MongoDB database with the provided JSON data
print('GO V RADUGU');

const fs = require('fs');

// Read the JSON file
const data = JSON.parse(fs.readFileSync('/records.json', 'utf8'));

// Connect to the database and insert the data
db = connect("mongodb://localhost:27017/lab4");


const convertUser = (users) => {
    return users.map(user => {
        user._id = ObjectId(user._id);
        user.basket_id = ObjectId(user.basket_id);
        user.creationDate = ISODate(user.creationDate);
        return user;
    });
};

const convertProduct = (products) => {
    return products.map(product => {
        product._id = ObjectId(product._id);
        return product;
    });
};

const convertBasket = (baskets) => {
    return baskets.map(basket => {
        basket._id = ObjectId(basket._id);
        basket.userId = ObjectId(basket.userId);
        basket.products = convertProduct(basket.products);
        return basket;
    });
};

db.users.insertMany(convertUser(data.users));
db.baskets.insertMany(convertBasket(data.baskets));
db.products.insertMany(convertProduct(data.products));