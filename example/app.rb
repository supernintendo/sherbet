#!/usr/bin/env ruby

require 'sinatra'

set :bind, '0.0.0.0'

get '/' do
  <<-eos
    <!doctype html>
    <html>
    <head>
        <title>test page</title>
    </head>
    <body>
        <div id="hello-world">
            Hello world!
        </div>
    </body>
    </html>
  eos
end
