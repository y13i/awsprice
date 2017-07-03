require "formula"

class Awsprice < Formula
  VERSION = "0.0.3"

  homepage "https://github.com/y13i/awsprice"
  version  VERSION
  url      "https://github.com/y13i/awsprice/releases/download/v#{VERSION}/awsprice-darwin_amd64.tar.gz"
  sha256   "80d322b0bacfbc106735021e38075e222dee80d1951f9f570858e268d41bc272"

  head "https://github.com/y13i/awsprice.git", :branch => "master"

  def install
    bin.install "awsprice"
  end
end
