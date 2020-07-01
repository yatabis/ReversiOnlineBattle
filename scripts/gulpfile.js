const pug          = require("gulp-pug");
const sass         = require("gulp-sass");
const postcss      = require("gulp-postcss");
const prefixer = require("autoprefixer");
const babel        = require("gulp-babel");
const uglify       = require("gulp-uglify-es").default;

const {task, src, dest, parallel} = require("gulp");

// Paths
const paths = {
  "htmlSrc" : "../website/src/pug/*.pug",
  "cssSrc"  : "../website/src/sass/*.sass",
  "jsSrc"   : "../website/src/js/**/*.js",
  "htmlDist": "../website/static/html/",
  "cssDist" : "../website/static/css/",
  "jsDist"  : "../website/static/js/"
}

//Pug
task("pug", () => {
  return src(paths.htmlSrc)
      .pipe(pug({
        pretty: false
      }))
      .pipe(dest(paths.htmlDist));
});

// Sass
task("sass", function () {
  return (
    src(paths.cssSrc)
    .pipe(sass({
      outputStyle: "compressed"
    }))
    .pipe(postcss([prefixer({
      cascade: false,
      grid: true
    })]))
    .pipe(dest(paths.cssDist))
  );
});

//JS Compress
task("js", function () {
  return (
    src(paths.jsSrc)
    .pipe(babel())
    .pipe(uglify({}))
    .pipe(dest(paths.jsDist))
  );
});

task("all", parallel("pug", "sass", "js"))
