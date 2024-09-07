// API endpoint URL
const apiEndpoint = 'http://localhost:8081/api'; // replace with your API endpoint

// get books button
const showBooksBtn = document.getElementById('show-books-btn');

// main container
const mainContainer = document.getElementById('main-container');

// event listener for show books button
showBooksBtn.addEventListener('click', () => {
    // fetch books from API
    fetch(`${apiEndpoint}/books`)
        .then(response => response.json())
        .then(response => {
            // clear main container
            mainContainer.innerHTML = '';

            // create book containers
            response.books.forEach(book => {
                const bookContainer = document.createElement('div');
                bookContainer.classList.add('book-container');

                const title = document.createElement('h2');
                title.textContent = book.title;

                const author = document.createElement('p');
                author.textContent = book.author;

                const showQuotesBtn = document.createElement('button');
                showQuotesBtn.textContent = 'Show Quotes';
                showQuotesBtn.addEventListener('click', () => {
                    // fetch quotes for book from API
                    fetch(`${apiEndpoint}/quotes/${book.id}`)
                        .then(response => response.json())
                        .then(response => {
                            // create quote containers
                            const quoteContainer = document.createElement('div');
                            quoteContainer.classList.add('quote-container');

                            response.quotes.forEach(quote => {
                                const blockquote = document.createElement('blockquote');
                                const paragraph = document.createElement('p');
                                paragraph.textContent = quote.text;

                                blockquote.appendChild(paragraph);
                                quoteContainer.appendChild(blockquote);
                            });

                            // append quote container to book container
                            bookContainer.appendChild(quoteContainer);
                        })
                        .catch(error => console.error(error));
                });

                bookContainer.appendChild(title);
                bookContainer.appendChild(author);
                bookContainer.appendChild(showQuotesBtn);

                mainContainer.appendChild(bookContainer);
            });
        })
        .catch(error => console.error(error));
});