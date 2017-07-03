require "formula"

class Awsprice < Formula
  VERSION = "0.0.2"

  homepage "https://github.com/y13i/awsprice"
  version  VERSION
  url      "https://github.com/y13i/awsprice/releases/download/v#{VERSION}/awsprice-darwin_amd64.tar.gz"
  sha256   "d1621b5b7c160da47c216e77c1bc066ba9591de78eb79ddc4441330bbceb6f80"

  head "https://github.com/y13i/awsprice.git", :branch => "master"

  def install
    bin.install "awsprice"
  end
end
