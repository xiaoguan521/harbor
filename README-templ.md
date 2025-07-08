# Harbor Templ UI

This directory contains the new templ-based UI for Harbor, replacing the previous Angular frontend.

## Overview

The templ UI provides a modern, fast, and maintainable frontend for Harbor using:

- **[a-h/templ](https://github.com/a-h/templ)**: Type-safe HTML templating in Go
- **Native Go**: Server-side rendering with Go's built-in HTTP server
- **Minimal Dependencies**: No complex frontend build pipeline
- **Type Safety**: Compile-time template validation

## Structure

```
templates/
├── base.templ           # Base HTML layout
├── navigation.templ     # Top navigation bar
├── sidenav.templ       # Side navigation menu
├── dashboard.templ     # Main dashboard page
└── *_templ.go          # Generated Go files (auto-generated)

cmd/templ-server/
└── main.go             # HTTP server for templ UI
```

## Quick Start

### Prerequisites

- Go 1.21 or later
- templ CLI tool (installed automatically)

### Development

1. **Generate templates:**
   ```bash
   make -f Makefile.templ templ-generate
   ```

2. **Run development server:**
   ```bash
   make -f Makefile.templ templ-dev
   ```

3. **Visit the UI:**
   Open [http://localhost:8080](http://localhost:8080) in your browser

### Available Make Targets

- `templ-generate`: Generate Go code from templ templates
- `templ-server`: Build and run the server
- `templ-dev`: Run in development mode
- `templ-build`: Build server binary
- `templ-clean`: Clean generated files
- `templ-install`: Install templ CLI tool

## Features Implemented

### ✅ Core UI Components

- **Base Layout**: HTML5 template with Harbor branding
- **Navigation**: Top navigation bar with user menu
- **Side Navigation**: Hierarchical menu structure
- **Dashboard**: Main landing page with cards

### ✅ Styling

- **Harbor Colors**: Uses official Harbor color scheme
- **Responsive Design**: Works on desktop and mobile
- **Modern CSS**: Flexbox layout, hover effects, transitions

### ✅ Navigation Structure

- **Projects**: Main project listing (placeholder)
- **Logs**: System logs access
- **System Management**: Admin functions
  - Users management
  - Groups management
  - Configuration settings

## Migration from Angular

### What's Different

| Angular | Templ |
|---------|-------|
| Client-side rendering | Server-side rendering |
| TypeScript | Go types |
| Complex build pipeline | Simple `go build` |
| npm/node_modules | Go modules |
| Angular CLI | templ CLI |

### Migration Benefits

1. **Performance**: Server-side rendering is faster
2. **Type Safety**: Compile-time template validation
3. **Simplicity**: No complex frontend tooling
4. **Maintainability**: Go developers can work on UI
5. **Security**: Server-side rendering reduces XSS risks

## Development Workflow

### 1. Edit Templates

Edit `.templ` files in the `templates/` directory:

```go
// templates/mypage.templ
package templates

type MyPageProps struct {
    Title string
    Items []string
}

templ MyPage(props MyPageProps) {
    @Base(BaseProps{
        Title: props.Title,
        Content: myPageContent(props),
    })
}

templ myPageContent(props MyPageProps) {
    <div class="my-page">
        <h1>{props.Title}</h1>
        <ul>
            for _, item := range props.Items {
                <li>{item}</li>
            }
        </ul>
    </div>
}
```

### 2. Generate Go Code

```bash
make -f Makefile.templ templ-generate
```

### 3. Add Route Handler

```go
// cmd/templ-server/main.go
func myPageHandler(w http.ResponseWriter, r *http.Request) {
    props := templates.MyPageProps{
        Title: "My Page",
        Items: []string{"Item 1", "Item 2"},
    }
    
    ctx := context.Background()
    if err := templates.MyPage(props).Render(ctx, w); err != nil {
        http.Error(w, "Error", http.StatusInternalServerError)
        return
    }
}

// In main()
http.HandleFunc("/mypage", myPageHandler)
```

### 4. Test

```bash
make -f Makefile.templ templ-dev
```

## Integration with Harbor Backend

The templ UI integrates with Harbor's existing Go backend:

1. **API Calls**: Use Harbor's existing REST APIs
2. **Authentication**: Leverage Harbor's session management
3. **Configuration**: Use Harbor's configuration system
4. **Logging**: Use Harbor's logging framework

## Production Deployment

### Build

```bash
make -f Makefile.templ templ-build
```

### Deploy

The `bin/templ-server` binary can be deployed alongside Harbor's main services.

## Future Enhancements

- [ ] Complete project management UI
- [ ] Repository and artifact views
- [ ] User management interface
- [ ] Configuration management
- [ ] Real-time updates with WebSockets
- [ ] Internationalization (i18n)
- [ ] Advanced search and filtering
- [ ] Integration with Harbor's existing middleware

## Contributing

1. Follow Go coding standards
2. Use `gofmt` for code formatting
3. Generate templates after changes: `make -f Makefile.templ templ-generate`
4. Test changes with: `make -f Makefile.templ templ-dev`
5. Ensure type safety - templates are compiled with Go

## Resources

- [templ Documentation](https://templ.guide/)
- [Harbor API Documentation](https://goharbor.io/docs/)
- [Go HTML Templates](https://pkg.go.dev/html/template)