class Iq < Formula
  desc "INI file query tool"
  homepage "https://github.com/djui/iq"
  url "https://github.com/djui/iq/archive/v1.0.0.tar.gz"
  sha256 "3ddc16325f44f3b5aeeaf172d2462a31ca8d85bbe004eef161bad620262078ec"

  bottle do
    cellar :any_skip_relocation
    rebuild 1
    sha256 "c044b182bc40f0ca86ef5eda4ec045836d1d2c045699e31104837a48c29d654d" => :mojave
    sha256 "c044b182bc40f0ca86ef5eda4ec045836d1d2c045699e31104837a48c29d654d" => :high_sierra
    sha256 "c044b182bc40f0ca86ef5eda4ec045836d1d2c045699e31104837a48c29d654d" => :sierra
    sha256 "c044b182bc40f0ca86ef5eda4ec045836d1d2c045699e31104837a48c29d654d" => :el_capitan
  end

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = HOMEBREW_CACHE/"go_cache"
    (buildpath/"src/github.com/djui/iq").install buildpath.children

    cd "src/github.com/djui/iq" do
      system "go", "build", "-o", bin/"iq", "main.go"
    end
  end

  test do
    system "#{bin}/iq", "-help"
  end
end
