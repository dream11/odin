package auth

const (
	successPageTemplate = `<!DOCTYPE html>
    <html>
    <head>
        <meta charset="utf-8">
        <title>Continue in Terminal</title>
    </head>
    <body style="font-family:Segoe UI,Arial,sans-serif;text-align:center;padding-top:10%;background:#f5f9ff;">
        <div style="display:inline-block;background:#fff;border-radius:10px;padding:30px 40px;box-shadow:0 4px 12px rgba(0,0,0,0.1);">
            <div style="font-size:40px;color:#1a73e8;"></div>
            <h2 style="color:#1a73e8;margin:10px 0;">Authentication received</h2>
            <p style="color:#555;">Completing sign-in in your terminal</p>
        </div>
    </body>
    </html>`

	errorPageTemplate = `<!DOCTYPE html><html><head><meta charset="utf-8"><title>%s</title></head><body style="font-family:Segoe UI,Arial,sans-serif;text-align:center;padding-top:10%%;"><h2>✗ %s</h2><p>%s</p></body></html>`
)

