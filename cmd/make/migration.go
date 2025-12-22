package make

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Migration(name string) {
	if len(os.Args) < 2 {
		fmt.Println("Pake: go run make:migration NamaTable")
		return
	}

	// name := os.Args[1]
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "_"))
	ts := time.Now().Format("20060102150405")

	base := fmt.Sprintf("%s_%s", ts, slug)
	up := base + ".up.sql"
	down := base + ".down.sql"

	dir := "db/migrations"
	os.MkdirAll(dir, 0o755)

	upPath := filepath.Join(dir, up)
	downPath := filepath.Join(dir, down)

	os.WriteFile(upPath, []byte(fmt.Sprintf(`	
CREATE TABLE IF NOT EXISTS %s (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

	-- write up migration here
);

-- Trigger agar updated_at auto-update
CREATE OR REPLACE FUNCTION update_%s_updated_at()
RETURNS TRIGGER AS $$
BEGIN
	NEW.updated_at = NOW();
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER %s_updated_at_trigger
BEFORE UPDATE ON %s
FOR EACH ROW
EXECUTE FUNCTION update_%s_updated_at();
`, name, name, name, name, name)), 0o644)

	os.WriteFile(downPath, []byte(fmt.Sprintf(`
-- YANG SALAH (akan error jika tabel tidak ada):
-- DROP TABLE %s;

-- YANG BENAR (aman meski tabel sudah tidak ada):
DROP TABLE IF EXISTS %s;

-- Atau jika ada foreign key constraints:
-- DROP TABLE IF EXISTS %s CASCADE;
`, name, name, name)), 0o644)

	fmt.Println("Created:", upPath)
	fmt.Println("Created:", downPath)
}
